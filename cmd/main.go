package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 15432
	dbUser     = "postgres"
	dbPassword = "pass001"
	dbName     = "postgres"
)

var db *sql.DB

func main() {
	// PostgreSQLデータベースに接続
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()

	// ルーティングの設定
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", handleLogin)

	r.Run(":8080")
}

func handleLogin(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSONデータの解析に失敗しました"})
		return
	}

	username := request.Username
	password := request.Password

	// ユーザー情報をデータベースから取得
	storedPassword := "0000"
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "ユーザーが存在しません"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "エラーが発生しました"})
		}
		return
	}

	// パスワードを検証
	// err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if password != storedPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "パスワードが正しくありません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ログイン成功"})
}
