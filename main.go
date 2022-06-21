package main

import (
	"github.com/Lemon-Potato/SingShark/global"
	"github.com/Lemon-Potato/SingShark/pkg/logger"
	"github.com/Lemon-Potato/SingShark/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(global.ServerSetting.HttpPort)
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Logger", &global.LoggerSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	logger, err := logger.NewLogger(
		global.LoggerSetting.FileName,
		global.LoggerSetting.MaxSize,
		global.LoggerSetting.MaxBackup,
		global.LoggerSetting.MaxAge,
		global.LoggerSetting.Compress,
		global.LoggerSetting.Type,
		global.LoggerSetting.Level,
	)
	if err != nil {
		return err
	}
	global.Logger = logger
	return nil
}
