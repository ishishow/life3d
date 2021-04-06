package postgre

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const driverName = "postgres"

type SQLHandler struct {
	Conn *sql.DB
}

func NewSQLHandler() SQLHandler {
	// config
	//user := os.Getenv("POSTGRES_USER")
	//password := os.Getenv("POSTGRES_PASSWORD")
	//host := os.Getenv("POSTGRES_HOST")
	//port := os.Getenv("POSTGRES_PORT")
	//database := os.Getenv("POSTGRES_DB")
	databaseUrl := os.Getenv("DATABASE_URL")

	conn, err := sql.Open(driverName, databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("connect!")
	}

	return SQLHandler{
		Conn: conn,
	}
}
