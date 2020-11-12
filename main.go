package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	DBconf := NewDBconfig()

	dbURL := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", DBconf.host, DBconf.port, DBconf.user, DBconf.pass, DBconf.dbname)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Connection failed")
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Print("db doesnt ping")
	}

}
