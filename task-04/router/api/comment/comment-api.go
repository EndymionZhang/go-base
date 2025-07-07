package comment

import (
	"github.com/endymion/go-base/task-04/common/constant/responseCode"
	"github.com/endymion/go-base/task-04/common/model/request"
	"github.com/endymion/go-base/task-04/common/model/response"
	"github.com/endymion/go-base/task-04/model"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	value, _ := c.Get("username")
	username := value.(string)
	user, _ := model.FindUserByName(username)
	var req request.CommentCreateRequest
	code := request.Request(c, request.Post, &req)
	if code != responseCode.Success {
		response.Fail(code, c)
		return
	}
	comment := model.Comment{
		Content: req.Content,
		PostID:  req.PostID,
		UserId:  user.ID,
	}
	err := comment.CreateComment()
	if err != nil {
		response.FailWithDetailed(responseCode.Error, response.Response{}, c)
		return
	}
	response.Ok(c)
}

func ListComments(c *gin.Context) {
	var req request.IdRequest
	code := request.Request(c, request.Get, &req)
	if code != responseCode.Success {
		response.Fail(code, c)
		return
	}
	comments := model.ListComments(req.Id)
	response.OkWithData(comments, c)
}
