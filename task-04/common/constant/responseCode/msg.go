package responseCode

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",

	ErrorAuthCheckTokenFail: "Token鉴权失败，请重新登陆",
	ErrorUserNotExist:       "用户不存在",
	ErrorPasswordWrong:      "密码错误",
	ErrorDeletePost:         "删除文章失败",
	ErrorEditPost:           "修改文章失败",

	ErrorUserExist: "用户已存在",

	SystemError: "系统内部错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
