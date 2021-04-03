package postgre

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const driverName = "postgre"

type SQLHandler struct {
	Conn *sql.DB
}

func NewSQLHandler() SQLHandler {
	// config
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	database := os.Getenv("POSTGRES_DATABASE")

	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database
	conn, err := sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}

	return SQLHandler{
		Conn: conn,
	}
}
