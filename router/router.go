package router

import (
	"time"

	"go_project/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// New 创建并配置 HTTP 路由
func New() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization", "Token"},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          12 * time.Hour,
	}))
	controllers.RegisterRoutes(r)
	return r
}
