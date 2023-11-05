package api

import (
	"app1-be/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func SetupRoutes(router *gin.Engine, dbInfo string) {

	// GET:/api/hello
	router.GET("/api/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	// POST:/api/sendUserInfo
	router.POST("/api/sendUserInfo", func(c *gin.Context) {
		service.SendUserInfoService(c, dbInfo)
	})

}
