package main

import (
	"app1-be/internal/api"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	api.SetupRoutes(router)

	router.Run(":8080")
}
