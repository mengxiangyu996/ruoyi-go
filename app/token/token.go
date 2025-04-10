package token

import (
	"context"
	"encoding/json"
	"errors"
	"ruoyi-go/app/dto"
	rediskey "ruoyi-go/common/types/redis-key"
	"ruoyi-go/common/uuid"
	"ruoyi-go/config"
	"ruoyi-go/framework/dal"
	"ruoyi-go/framework/datetime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// 授权声明
type SysUserClaim struct {
	Uuid string `json:"uuid"`
	jwt.RegisteredClaims
}

// 获取授权声明
func GetClaims() *SysUserClaim {

	uuid, _ := uuid.New()

	return &SysUserClaim{
		Uuid: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
			Issuer:    "ruoyi-go",                     // 签发人
		},
	}
}

// 生成token
func (a *SysUserClaim) GenerateToken(user dto.UserTokenResponse) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, a).SignedString([]byte(config.Data.Token.Secret))
	if err != nil {
		return "", err
	}

	err = dal.Redis.Set(context.Background(), rediskey.UserTokenKey+a.Uuid, &UserTokenResponse{
		UserTokenResponse: user,
		ExpireTime:        datetime.Datetime{Time: time.Now().Add(time.Minute * time.Duration(config.Data.Token.ExpireTime))},
	}, time.Minute*time.Duration(config.Data.Token.ExpireTime)).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}

// 刷新token
func RefreshToken(ctx *gin.Context, user dto.UserTokenResponse) {

	tokenKey, err := getUserTokenKey(ctx)
	if err != nil {
		return
	}

	dal.Redis.Set(ctx.Request.Context(), tokenKey, &UserTokenResponse{
		UserTokenResponse: user,
		ExpireTime:        datetime.Datetime{Time: time.Now().Add(time.Minute * time.Duration(config.Data.Token.ExpireTime))},
	}, time.Minute*time.Duration(config.Data.Token.ExpireTime))
}

// 解析token
func GetAuhtUser(ctx *gin.Context) (*UserTokenResponse, error) {

	tokenKey, err := getUserTokenKey(ctx)
	if err != nil {
		return nil, err
	}

	var user UserTokenResponse

	if err = dal.Redis.Get(ctx.Request.Context(), tokenKey).Scan(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// 删除token
func DeleteToken(ctx *gin.Context) error {

	tokenKey, err := getUserTokenKey(ctx)
	if err != nil {
		return err
	}

	return dal.Redis.Del(ctx.Request.Context(), tokenKey).Err()
}

// 获取授权用户的redis key
func getUserTokenKey(ctx *gin.Context) (string, error) {

	authorization := ctx.GetHeader(config.Data.Token.Header)
	if authorization == "" {
		return "", errors.New("请先登录")
	}

	tokenSplit := strings.Split(authorization, " ")
	if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
		return "", errors.New("authorization format error")
	}

	token, err := jwt.ParseWithClaims(tokenSplit[1], &SysUserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Data.Token.Secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", errors.New("token格式错误")
			}
			if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", errors.New("token未生效")
			}
			return "", errors.New("token校验失败")
		}
		return "", err
	}

	if claims, ok := token.Claims.(*SysUserClaim); ok && token.Valid {
		return rediskey.UserTokenKey + claims.Uuid, nil
	}

	return "", errors.New("token校验失败")
}

type UserTokenResponse struct {
	dto.UserTokenResponse
	ExpireTime datetime.Datetime `json:"expireTime"`
}

// 序列化dto.UserTokenResponse，实现redis读写
func (u UserTokenResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// 反序列化dto.UserTokenResponse，实现redis读写
func (u *UserTokenResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
