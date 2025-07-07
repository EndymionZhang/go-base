package response

import (
	"github.com/endymion/go-base/task-04/common/constant/responseCode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	*gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		responseCode.GetMsg(code),
	})
}

func Ok(c *gin.Context) {
	Result(responseCode.Success, map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(responseCode.Success, data, c)
}

func Fail(responseCode int, c *gin.Context) {
	Result(responseCode, map[string]interface{}{}, c)
}

func NoAuth(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		responseCode.ErrorAuthCheckTokenFail,
		nil,
		responseCode.GetMsg(responseCode.ErrorAuthCheckTokenFail),
	})
}

func FailWithDetailed(responseCode int, data interface{}, c *gin.Context) {
	Result(responseCode, data, c)
}
