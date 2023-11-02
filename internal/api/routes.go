package api

import (
	"app1-be/internal/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func SetupRoutes(router *gin.Engine) {

	// CORSの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Next.jsのサーバーのオリジンを許可
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// PostgreSQL への接続情報を設定
	const (
		host     = "localhost"
		port     = 15432
		user     = "yourusername"
		password = "yourpassword"
		dbname   = "yourdbname"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// GETリクエストをハンドルするエンドポイント
	router.GET("/api/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	// /api/sendUserInfo でリクエストを受け取るエンドポイント
	router.POST("/api/sendUserInfo", func(c *gin.Context) {
		var newUser model.User
		// リクエストボディからユーザー情報を取得
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// データベースに新しいユーザーを登録
		sqlStatement := `
     INSERT INTO users (name, email)
     VALUES ($1, $2)
     RETURNING id`
		id := 0
		err = db.QueryRow(sqlStatement, newUser.Name, newUser.Email).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 成功レスポンスを送信
		c.JSON(http.StatusOK, gin.H{"status": "user created", "id": id})
	})

}