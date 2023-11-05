package main

import (
	"app1-be/internal/api"
	"app1-be/internal/common"
	"app1-be/internal/test"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	// CORSの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// DB接続情報の取得
	dbInfo := common.GetDbInfo()

	// テスト用DB作成
	test.SetupTestDb(dbInfo)

	// ルーティング呼び出し
	api.SetupRoutes(router, dbInfo)

	// 8080ポートで待ち受け開始
	router.Run(":8080")
}
