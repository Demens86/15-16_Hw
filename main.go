package main

import (
	"fmt"
	"log"
	"net/http"
	"prac/internal/handlers"
	"prac/internal/repo"
	"prac/internal/usecase"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	//fmt.Println("Start")

	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Println("db connection success")

	repo := repo.NewRepo(db)
	uc := usecase.NewCase(repo)
	h := handlers.NewHandler(uc)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", h.CreateUser)

	//fmt.Println("End")

	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
