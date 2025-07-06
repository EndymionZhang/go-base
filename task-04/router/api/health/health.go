package health

import (
	"github.com/endymion/go-base/task-04/common/model/response"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	response.Ok(c)
}
