package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectarDB() *sql.DB {
	// user-database-password-host-ssl
	conexao := "user=postgres dbname=alura_loja password=0512 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		// Encerra a aplicação
		panic(err.Error())
	}

	return db
}
