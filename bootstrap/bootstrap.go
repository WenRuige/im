package bootstrap

import (
	"context"
	"database/sql"
	"time"

	"github.com/didi/gendry/scanner"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/im/library/resource"
	"go.uber.org/zap"
)

func MustLoadAppConfig() {

}

func MustInit(ctx context.Context) {

	// 1.初始化MYSQL
	InitMySQL()

	// 2.初始化Rdis

	// 3.初始化zapper
	InitLogger()

}

func InitLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	resource.Logger = logger
}

// InitMySQL 初始化MySQL
func InitMySQL() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "123456",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "im",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// 设置scanner
	scanner.SetTagName("db")

	resource.IMDB = db
}
