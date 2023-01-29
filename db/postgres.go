package db

import (
	"database/sql"
	"log"
	"sync"

	"escort-book-user-log-consumer/config"

	_ "github.com/lib/pq"
)

var postgresInstance *PostgresClient
var singlePostgresClient sync.Once

type PostgresClient struct {
	UserDB *sql.DB
}

func initPostgresClient() {
	userDB, err := sql.Open(
		"postgres",
		config.InitializePostgres().Databases.User,
	)

	if err != nil {
		log.Fatalf("Error connecting with user DB: %s", err.Error())
	}

	postgresInstance = &PostgresClient{
		UserDB: userDB,
	}
}

func NewPostgresClient() *PostgresClient {
	singlePostgresClient.Do(initPostgresClient)
	return postgresInstance
}
