package statuscode

// 状态码
var (
	Success      = 200 // 成功
	Error        = 500 // 失败
	BadRequest   = 400 // 参数列表错误
	Unauthorized = 401 // 未授权
)

var statusMessage = map[int]string{
	Success:      "成功",
	Error:        "失败",
	BadRequest:   "参数列表错误",
	Unauthorized: "未授权",
}

func GetMessage(code int) string {

	if msg, ok := statusMessage[code]; ok {
		return msg
	}

	return "未知错误"
}
