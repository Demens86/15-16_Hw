package models

import "time"

//var User []User

type User struct {
	Id        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Age       int       `db:"age" json:"age"`
	IsActive  bool      `db:"is_active" json:"is_active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type UserRequest struct {
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Age      int    `db:"age" json:"age"`
	IsActive bool   `db:"is_active" json:"is_active"`
}
