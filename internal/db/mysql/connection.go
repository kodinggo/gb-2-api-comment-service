package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kodinggo/gb-2-api-comment-service/internal/config"
)

func NewConnStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		config.MySQLDBUser(),
		config.MySQLDBPass(),
		config.MySQLDBHost(),
		config.MySQLDBName())
}

func InitDBConn() *sql.DB {
	dbConn, err := sql.Open("mysql", NewConnStr())
	if err != nil {
		panic(err)
	}

	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)

	if err := dbConn.Ping(); err != nil {
		panic(err)
	}

	return dbConn
}
