package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func main() {
	fmt.Println("Start")

	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@localhost:5432/db_prac?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Println("db connection success")

	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("End")
}
