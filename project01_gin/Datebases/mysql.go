package Datebases

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func init() {
	// 数据库连接字符串
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo1_database")
	if err != nil {
		log.Fatal(err)
	}

	// 测试数据库连接
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
}
