package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"

	"dev.go/app"
	"dev.go/databases"
	"github.com/go-redis/redis"
)

var tpl *template.Template
var err error
var webapp app.App

func init() {

	// Redis setup for session management
	databases.SessionDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	webapp = app.App{}

	// MySql database Configuration
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               os.Getenv("DBPASSWD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:13306",
		DBName:               "testDB",
		AllowNativePasswords: true,
	}

	// MySql database initialization
	databases.DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}

	tpl = template.Must(template.ParseGlob("assets/templates/*"))
}

func main() {
	webapp.Init(*tpl)

	// Testing connection.

	err = databases.DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	err = databases.SessionDB.Ping().Err()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connection secured")

	webapp.Start()

}
