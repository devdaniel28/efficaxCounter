package models

import "time"

type User struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	CreatAt time.Time `json:"created_at"`
}