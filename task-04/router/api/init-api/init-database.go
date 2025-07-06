package init_api

import (
	"github.com/endymion/go-base/task-04/common/constant/responseCode"
	"github.com/endymion/go-base/task-04/common/model/response"
	"github.com/endymion/go-base/task-04/model"
	"github.com/gin-gonic/gin"
)

func InitDb(c *gin.Context) {
	err := model.InitDb()
	if err != nil {
		response.Fail(responseCode.Error, c)
		return
	}
}
