package response

import (
	statusCode "ruoyi-go/common/types/status-code"

	"github.com/gin-gonic/gin"
)

// 响应
type Response struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"-"`
}

type OperationType int

// 初始化成功响应
func NewSuccess() *Response {

	return &Response{
		Code: statusCode.Success,
		Msg:  statusCode.GetMessage(statusCode.Success),
		Data: make(map[string]interface{}),
	}
}

// 初始化失败响应
func NewError() *Response {

	return &Response{
		Code: statusCode.Error,
		Msg:  statusCode.GetMessage(statusCode.Error),
		Data: make(map[string]interface{}),
	}
}

// 设置响应码
func (r *Response) SetCode(code int) *Response {

	r.Code = code

	r.Msg = statusCode.GetMessage(code)

	return r
}

// 设置响应信息
func (r *Response) SetMsg(msg string) *Response {

	if msg == "" {
		r.Msg = statusCode.GetMessage(r.Code)
		return r
	}

	r.Msg = msg
	return r
}

// 设置响应数据
func (r *Response) SetData(key string, value interface{}) *Response {

	if key == "code" || key == "msg" {
		return r
	}

	r.Data[key] = value

	return r
}

// 设置分页响应数据
func (r *Response) SetPageData(rows interface{}, total int) *Response {

	r.Data["rows"] = rows
	r.Data["total"] = total

	return r
}

// 设置响应数据
func (r *Response) SetDataMap(data map[string]interface{}) *Response {

	for key, value := range data {
		if key == "code" || key == "msg" {
			continue
		}
		r.Data[key] = value
	}

	return r
}

// 序列化返回
func (r *Response) ToJson(ctx *gin.Context) {

	response := gin.H{
		"code": r.Code,
		"msg":  r.Msg,
	}

	for key, value := range r.Data {
		response[key] = value
	}

	ctx.JSON(200, response)
}
