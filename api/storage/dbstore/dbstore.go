package dbstore

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/storage/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"reflect"
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

func (db *DbConnection) GetPosts(request *storage.QueryRequest) []*models.Post {
	panic("Implement me")
}

func (db *DbConnection) CreatePost(post *models.Post) (bool, error) {
	panic("Implement me")
}

func (db *DbConnection) Stop() {
	db.Connection.Close()
}

func (db *DbConnection) RunMigrations() error {
	ms := db.getExistedTables()

	//TODO: make as transaction
	for _, model := range ms {
		if m, ok := model.(models.Migration); ok {
			sql := m.TableCreate()

			_, err := db.Connection.Exec(sql)
			if err != nil {
				return fmt.Errorf("cannot make migration: %s, %v", sql, err)
			}

			log.Printf("[DBSTORAGE] migrations for: %s SUCCESSFUL", db.getType(m))
		}
	}

	return nil
}

func (db *DbConnection) getType(strct interface{}) string {
	if t := reflect.TypeOf(strct); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (db *DbConnection) getExistedTables() []interface{} {
	//TODO create a query and get all migrations
	return []interface{}{
		&models.MigrationModel{},
		&models.Tag{},
		&models.Post{},
		&models.Category{},
	}
}
