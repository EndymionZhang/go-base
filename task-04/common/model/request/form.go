package request

import (
	"github.com/astaxie/beego/validation"
	"github.com/endymion/go-base/task-04/common/constant/responseCode"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, responseCode.InvalidParams
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, responseCode.Error
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, responseCode.InvalidParams
	}

	return http.StatusOK, responseCode.Success
}
