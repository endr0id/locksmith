package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=local_user password=password dbname=auth_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic("DB接続失敗: " + err.Error())
	}

	fmt.Println("DB接続成功")
}