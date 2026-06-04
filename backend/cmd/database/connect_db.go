package database

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	dbURL := os.Getenv("DBURL")

	if dbURL == "" {
		return nil, fmt.Errorf("Error: Env database url null \n")
	}
 
	var err error

	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("Erro ao conectar ao banco de dados: %w \n", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("Error: Ping error, database closed: %w \n", err)
		
	}

	return db, nil
}
