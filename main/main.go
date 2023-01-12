package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"html"
	"log"
	"net/http"
)

//globals values

var (
	Id       int64
	Name     string
	Price    float64
	Quantity int64
	Status   bool
)

var (
	server   = "."
	port     = 1433
	user     = ""
	password = ""
	database = "tienda"
)

var schema = `

    CREATE TABLE IF NOT EXISTS users (
	user_id    int,
  	first_name varchar(81),
  	last_name  varchar(80),
	email      varchar(250),
	password   varchar(250)
)

`

type User struct {
	UserID    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
	Password  sql.NullString
}

func main() {

	db, err := sqlx.Connect(
		"mssql",
		"sqlserver:;source=REIN801206-11;initial catalog=Tienda;trusted_connection=true",
	)

	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail;
	//db.MustExec(schema)

	//tx := db.MustBegin()
	/*tx.NamedExec(
		"INSERT INTO users (user_id, first_name, last_name, email) VALUES (:user_id, :first_name, :last_name, :email)",
		&User{UserID: 1, FirstName: "Jane", LastName: "Citizen", Email: "jane.citzen@example.com"},
	)

	tx.Commit()
	*/
	// Query the database, storing results in a []User (wrapped in []interface{})
	people := []User{}
	db.Select(&people, "SELECT * FROM users ORDER BY first_name ASC")

	fmt.Println(people)

	//Marshal the map
	b, _ := json.Marshal(people)

	//Prints the resulted json
	fmt.Printf("Marshalled data: %s\n", b)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		json.NewEncoder(w).Encode(people)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
