package router

import (
	"github.com/endymion/go-base/task-04/middleware/jwt"
	"github.com/endymion/go-base/task-04/router/api/health"
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
	userLoginGroup.GET("/jwt/check", user.JwtCheck)
	return Router
}
