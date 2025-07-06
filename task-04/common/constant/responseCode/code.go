package responseCode

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrorAuthCheckTokenFail = 10001
	ErrorUserNotExist       = 10002
	ErrorPasswordWrong      = 10003

	ErrorUserExist = 20001

	SystemError = 30001
)
