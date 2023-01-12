package dbConfig

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "."
	port     = 1433
	user     = ""
	password = ""
	database = "tienda"
)

func GetDB() (db *sql.DB, err error) {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)
	db, err = sql.Open("mssql", connectionString)
	return
}
