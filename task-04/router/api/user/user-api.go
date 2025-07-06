package user

import (
	"github.com/endymion/go-base/task-04/common/constant/responseCode"
	"github.com/endymion/go-base/task-04/common/logger"
	"github.com/endymion/go-base/task-04/common/model/request"
	"github.com/endymion/go-base/task-04/common/model/response"
	"github.com/endymion/go-base/task-04/common/util"
	"github.com/endymion/go-base/task-04/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	_, errorCode := request.BindAndValid(c, &user)
	if errorCode != responseCode.Success {
		logger.Error("参数验证错误, %s", user)
		response.FailWithDetailed(errorCode, nil, c)
		return
	}
	exist := model.IsUserNameExist(&model.User{Username: user.Username})
	if exist {
		logger.Error("用户已存在, %s", user)
		response.FailWithDetailed(responseCode.ErrorUserExist, nil, c)
		return
	}
	err := user.CreateUser()
	if err != nil {
		logger.Error("创建用户失败, %s", err)
		response.FailWithDetailed(responseCode.SystemError, nil, c)
		return
	}
	response.Ok(c)
}

func Login(c *gin.Context) {
	var user model.User
	_, err := request.BindAndValid(c, &user)
	if err != responseCode.Success {
		logger.Error("参数验证失败, %s", err)
		response.FailWithDetailed(err, nil, c)
		return
	}
	existUser, exist := model.FindUserByName(user.Username)
	if !exist {
		logger.Error("用户不存在, %s", err)
		response.FailWithDetailed(responseCode.ErrorUserNotExist, nil, c)
		return
	}
	if !util.ValidatePassword(user.Password, existUser.Password) {
		logger.Error("密码错误, %s", err)
		response.FailWithDetailed(responseCode.ErrorPasswordWrong, nil, c)
		return
	}
	token, _, error := util.GenerateToken(existUser.Username)
	if error != nil {
		logger.Error("生成token失败, %s", err)
		response.FailWithDetailed(responseCode.SystemError, nil, c)
		return
	}
	c.Header("x-jwt-token", token)
	response.OkWithData("登陆成功", c)
}

func JwtCheck(c *gin.Context) {
	username, _ := c.Get("username")
	response.OkWithData(username, c)
}
