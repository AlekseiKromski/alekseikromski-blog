package dbstore

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/storage/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"reflect"
	"strings"
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

// RunMigrations - take all migrations from table, compare and create tables, that should be
func (db *DbConnection) RunMigrations() error {
	notExisted, err := db.getNotExistedTables()
	if err != nil {
		return fmt.Errorf("trying to get all migrations: %w", err)
	}

	for _, model := range notExisted {
		if m, ok := model.(models.MigrationInterface); ok {
			sql := m.TableCreate()
			tableName := db.getType(m)

			_, err := db.Connection.Exec(sql)
			if err != nil {
				if strings.Contains(err.Error(), "already exists") {
					log.Printf("[DBSTORAGE] migrations for: %s ALREADY EXISTS", tableName)
					continue
				}
				return fmt.Errorf("cannot make migration: %s, %v", sql, err)
			}

			migration := models.CreateMigrationModel(tableName, sql)
			mq := migration.CreateRecord()
			_, err = db.Connection.Exec(mq)
			if err != nil {
				return fmt.Errorf("cannot create record in database %s: %w", mq, err)
			}

			log.Printf("[DBSTORAGE] Migration for: %s SUCCESSFUL", db.getType(m))
		}
	}

	return nil
}

// getType - get reflect type
func (db *DbConnection) getType(strct interface{}) string {
	if t := reflect.TypeOf(strct); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (db *DbConnection) getNotExistedTables() ([]models.MigrationInterface, error) {

	//All possible tables in application
	tables := []models.MigrationInterface{
		&models.Migration{},
		&models.Tag{},
		&models.Post{},
		&models.Category{},
	}

	query := `SELECT migrations."tableName" FROM migrations`
	rows, err := db.Connection.Query(query)
	if err != nil {
		if strings.Contains(err.Error(), "relation \"migrations\" does not exist") {
			return tables, nil
		}
		return nil, fmt.Errorf("cannot get all migrations from database: %w", err)
	}

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			return nil, fmt.Errorf("problem with scanning: %w", err)
		}

		model, index := containsInTables(tableName, tables)
		if model {
			tables = append(tables[:index], tables[index+1:]...)
		}
	}
	return tables, nil
}

// containsInTables - check if we have table name in list of tables, that implement migration interface
func containsInTables(tableName string, tables []models.MigrationInterface) (bool, int) {

	for index, table := range tables {
		if reflect.TypeOf(table).String() == fmt.Sprintf("*models.%s", tableName) {
			return true, index
		}
	}

	return false, 0
}
