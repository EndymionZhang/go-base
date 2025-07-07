package post

import (
	"github.com/endymion/go-base/task-04/common/constant/responseCode"
	"github.com/endymion/go-base/task-04/common/logger"
	"github.com/endymion/go-base/task-04/common/model/request"
	"github.com/endymion/go-base/task-04/common/model/response"
	"github.com/endymion/go-base/task-04/model"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	value, _ := c.Get("username")
	username := value.(string)
	user, _ := model.FindUserByName(username)
	var post model.Post
	code := request.Request(c, request.Post, &post)
	if code != responseCode.Success {
		response.Fail(code, c)
		return
	}
	post.AuthorID = user.ID
	err := post.CreatePost()
	if err != nil {
		logger.Error("创建文章失败, %s", err)
		response.Fail(responseCode.SystemError, c)
		return
	}
	response.Ok(c)
}

func ListPosts(c *gin.Context) {
	var searchRequest request.PostsSearchRequest
	code := request.Request(c, request.Get, &searchRequest)
	if code != responseCode.Success {
		response.Fail(code, c)
		return
	}
	count := model.CountPosts(&searchRequest)
	pageResponse := &response.PageResponse{
		Page:     searchRequest.Page,
		PageSize: searchRequest.PageSize,
		Total:    count,
	}
	if count == 0 {
		response.OkWithData(pageResponse, c)
		return
	}
	posts := model.ListPosts(&searchRequest)
	pageResponse.List = posts
	response.OkWithData(pageResponse, c)
}

func DeletePost(c *gin.Context) {
	var req request.IdRequest
	code := request.Request(c, request.Delete, &req)
	if code != responseCode.Success {
		response.Fail(code, c)
		return
	}
	value, _ := c.Get("username")
	username := value.(string)
	user, _ := model.FindUserByName(username)

	success, err := model.DeletePost(req.Id, user.ID)
	if err != nil {
		logger.Error("删除失败, %s", err)
		response.Fail(responseCode.SystemError, c)
		return
	}
	if !success {
		logger.Error("删除失败, %s", err)
		response.Fail(responseCode.ErrorDeletePost, c)
		return
	}
	response.Ok(c)
}

func GetPostDetail(c *gin.Context) {
	var req request.IdRequest
	code := request.Request(c, request.Get, &req)
	if code != responseCode.Success {
		response.FailWithDetailed(code, response.Response{}, c)
		return
	}
	post := model.GetPostDetail(req.Id)
	response.OkWithData(post, c)
	response.Ok(c)
}

func EditPost(c *gin.Context) {
	value, _ := c.Get("username")
	username := value.(string)
	user, _ := model.FindUserByName(username)
	var req model.Post
	code := request.Request(c, request.Put, &req)
	if code != responseCode.Success {
		response.FailWithDetailed(code, response.Response{}, c)
		return
	}
	if req.ID == 0 {
		response.FailWithDetailed(responseCode.InvalidParams, response.Response{}, c)
		return
	}
	success, err := model.EditPost(&req, user.ID)
	if !success || err != nil {
		response.FailWithDetailed(responseCode.ErrorEditPost, response.Response{}, c)
		return
	}
	response.OkWithData(response.Response{}, c)
}
