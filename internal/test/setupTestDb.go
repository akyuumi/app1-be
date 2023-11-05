package test

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func SetupTestDb(dbInfo string) {

	// データベースに接続
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブルを作成
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
            id serial PRIMARY KEY,
            name VARCHAR(255),
            email VARCHAR(255)
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	// テストデータを挿入
	_, err = db.Exec(`
        INSERT INTO users (name, email)
        VALUES ('naruke', 'naruke@example.com'), ('gouda', 'gouda@example.com'), ('iwadate', 'iwadate@example.com')
    `)
	if err != nil {
		log.Fatal(err)
	}
}
