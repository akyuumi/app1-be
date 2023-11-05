package service

import (
	"app1-be/internal/model"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func SendUserInfoService(c *gin.Context, psqlInfo string) {

	// DBオープン
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	// DB遅延クローズ
	defer db.Close()

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
}
