package util

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ErrCodeMsgMap = map[int]string{
	200: "登录成功",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	500: "Internal Server Error",
	502: "登录失败：用户名或者密码错误",
	504: "请求解析失败",
	505: "密码加密失败",
	506: "用户名或密码无效",
	507: "获取参数失败",
}

func NewError(errCode int) ErrorResponse {
	msg, ok := ErrCodeMsgMap[errCode]
	if !ok {
		msg = "Unknown Error"
	}
	return ErrorResponse{
		Code:    errCode,
		Message: msg,
	}
}
