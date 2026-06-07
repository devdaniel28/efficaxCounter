package models

import "database/sql"

type Counter struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Quote sql.NullInt64 `json:"quote"`
	UserId int `json:"userid"`
}