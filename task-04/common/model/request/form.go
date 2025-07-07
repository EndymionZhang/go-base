package request

import (
	"github.com/astaxie/beego/validation"
	"github.com/endymion/go-base/task-04/common/constant/responseCode"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) int {
	err := c.Bind(form)
	if err != nil {
		return responseCode.InvalidParams
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return responseCode.Error
	}
	if !check {
		MarkErrors(valid.Errors)
		return responseCode.InvalidParams
	}

	return responseCode.Success
}
