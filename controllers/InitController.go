package controllers

import (
	"github.com/gin-gonic/gin"
)

var R = gin.Default()

func InitController() {
	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode) // 或 gin.DebugMode
	// 定义路由处理函数
	R.GET("/", func(c *gin.Context) {
		c.String(200, "HTTPS服务运行在端口8002")
	})
	R.GET("/mysql/auto", mysqlAuto)
}
