package router

import (
	"go_project/config"
	"go_project/controllers"
	"time"

	"github.com/gin-contrib/cors"
)

func InitRouter() {
	controllers.R.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有域，生产环境建议改为具体域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	config.InitDB()
	controllers.InitController()
	err := controllers.R.Run(":8008")
	if err != nil {
		return
	}

	//StartHTTPS()
}
