package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"go-blog/pkg/logger"
	"time"
)

// DB 数据库对象
var DB *sql.DB

func Initialize() {
	initDB()
	createTables()
}

func initDB() {
	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "secret",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "go_blog",
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	//设置最大连接数
	DB.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	DB.SetMaxIdleConns(25)
	// 设置每个连接的过期时间
	DB.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接，失败报错
	err = DB.Ping()
	logger.LogError(err)
}

func createTables() {
	createArticlesSql := `CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci
	);`

	_, err := DB.Exec(createArticlesSql)
	logger.LogError(err)
}
