package models

import "time"

type User struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Password string   `json:"password"`
	CreatAt time.Time `json:"created_at"`
}