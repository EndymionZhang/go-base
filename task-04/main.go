package main

import (
	"fmt"
	"github.com/endymion/go-base/task-04/common/logger"
	"github.com/endymion/go-base/task-04/common/setting"
	"github.com/endymion/go-base/task-04/model"
	"github.com/endymion/go-base/task-04/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init() {
	setting.Load()
	logger.SetUp()
	model.SetUp()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := router.Routers()

	readTimeout := setting.ServerSetting.ReadTimeout * time.Second
	writeTimeout := setting.ServerSetting.WriteTimeout * time.Second
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	logger.Info("start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
