package main

import (
	"github.com/futhergo/websocketChat/internal/models"
	"github.com/futhergo/websocketChat/internal/pkg/DB"
	"github.com/futhergo/websocketChat/internal/pkg/settings"
	"github.com/futhergo/websocketChat/internal/router"
	logger "github.com/futhergo/websocketChat/log"
)

func main() {
	logger.Log2file("~~~~~~~~~~New Server Start~~~~~~~~~")
	settings.InitSetting()
	DB.InitDB()
	router.InitRoutes()

	if !DB.DB.HasTable(&models.UserEntity{}) {
		DB.DB.Create(&models.UserEntity{})
	}
	if !DB.DB.HasTable(&models.Message{}) {
		DB.DB.Create(&models.Message{})
	}
}