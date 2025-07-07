package router

import (
	"github.com/endymion/go-base/task-04/middleware/jwt"
	"github.com/endymion/go-base/task-04/router/api/comment"
	"github.com/endymion/go-base/task-04/router/api/health"
	"github.com/endymion/go-base/task-04/router/api/post"
	"github.com/endymion/go-base/task-04/router/api/user"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	Router.GET("health", health.Health)
	//Router.GET("init", init_api.InitDb)
	Router.POST("user/register", user.Register)
	Router.POST("user/login", user.Login)

	userLoginGroup := Router.Group("user").Use(jwt.JWT())
	{
		userLoginGroup.GET("/jwt/check", user.JwtCheck)
		userLoginGroup.POST("post/create", post.CreatePost)
		userLoginGroup.PUT("post/edit", post.EditPost)
		userLoginGroup.DELETE("post/delete", post.DeletePost)
		userLoginGroup.GET("post/list", post.ListPosts)

		userLoginGroup.POST("comment/create", comment.CreateComment)
		userLoginGroup.GET("comment/list", comment.ListComments)
	}
	return Router
}
