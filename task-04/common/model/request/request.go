package request

import (
	"github.com/astaxie/beego/validation"
	"github.com/endymion/go-base/task-04/common/constant/responseCode"
	"github.com/endymion/go-base/task-04/common/logger"
	"github.com/gin-gonic/gin"
)

// Request 请求处理
func Request(c *gin.Context, requestType RequestType, request interface{}) int {
	var err error
	switch requestType {
	case Post:
		err = c.ShouldBind(request)
	case Get:
		err = c.ShouldBindQuery(request)
	case Put:
		err = c.ShouldBindJSON(request)
	case Delete:
		err = c.ShouldBindJSON(request)
	default:
		return responseCode.InvalidParams
	}
	if err != nil {
		return responseCode.InvalidParams
	}
	valid := validation.Validation{}
	check, err := valid.Valid(request)
	if err != nil {
		return responseCode.Error
	}
	if !check {
		MarkErrors(valid.Errors)
		return responseCode.InvalidParams
	}
	return responseCode.Success
}

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logger.Info(err.Key, err.Message)
	}
	return
}
