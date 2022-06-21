package main

import (
	"github.com/Lemon-Potato/SingShark/global"
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
	return nil
}
