package captcha

import (
	"context"
	rediskey "ruoyi-go/common/types/redis-key"
	"ruoyi-go/framework/dal"
	"time"
)

// 实现captcha.Store接口
type RedisStore struct{}

func (r *RedisStore) Set(id string, value string) error {
	return dal.Redis.Set(context.Background(), rediskey.CaptchaCodeKey+id, value, time.Minute*5).Err()
}

func (r *RedisStore) Get(id string, clear bool) string {

	captcha, err := dal.Redis.Get(context.Background(), rediskey.CaptchaCodeKey+id).Result()
	if err != nil {
		return ""
	}

	if clear {
		if err = dal.Redis.Del(context.Background(), rediskey.CaptchaCodeKey+id).Err(); err != nil {
			return ""
		}
	}

	return captcha
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	return r.Get(id, clear) == answer
}
