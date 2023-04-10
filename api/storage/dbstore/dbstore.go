package dbstore

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DbConnection struct {
	Connection *sql.DB
}

func NewDbConnection(username, password, hostname, port, database string) (*DbConnection, error) {
	credits := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", hostname, port, username, password, database)
	db, err := sql.Open("postgres", credits)
	if err != nil {
		return nil, fmt.Errorf("cannot create database connection: %w", err)
	}
	log.Println("[DBSTORAGE] Pinging database")
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	log.Println("[DBSTORAGE] Connected to database")

	dbConnection := &DbConnection{
		Connection: db,
	}

	log.Println("[DBSTORAGE] Run migrations")
	if err = dbConnection.RunMigrations(); err != nil {
		return nil, fmt.Errorf("trying to run all migrations: %v", err)
	}

	return dbConnection, err
}

func (db *DbConnection) Stop() {
	db.Connection.Close()
}
