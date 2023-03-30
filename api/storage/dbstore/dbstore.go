package dbstore

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/storage/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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
	log.Println("[INFO] Pinging database")
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	log.Println("[INFO] Connected to database")
	return &DbConnection{
		Connection: db,
	}, nil
}

func (db *DbConnection) GetPosts(request *storage.QueryRequest) []*models.Post {
	panic("Implement me")
}

func (db *DbConnection) CreatePost(post *models.Post) (bool, error) {
	panic("Implement me")
}

func (db *DbConnection) Stop() {
	db.Connection.Close()
}
