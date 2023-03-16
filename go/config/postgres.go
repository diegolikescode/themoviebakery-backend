package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DisconnectPostgres func()

type PostgresConn struct {
	DbConn     sql.DB
	Disconnect DisconnectPostgres
}

func ConnectPostgres() /*RETURN ??*/ PostgresConn {
	connStr := "postgresql://admin:admin@localhost:5432/TheMovieBakery?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	newConnection := PostgresConn{
		DbConn: *db,
		Disconnect: func() {
			if err := db.Close(); err != nil {
				log.Fatal("erro disconnecting postgres")
			}
		},
	}

	return newConnection
}
