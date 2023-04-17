package dbstore

import (
	"alekseikromski.com/blog/api/storage/models"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type queueMigration struct {
	*models.TableCreation
	TableName string
	Model     models.MigrationInterface
}

// RunMigrations - take all migrations from table, compare and create tables, that should be
func (db *DbConnection) RunMigrations() error {
	notExisted, err := db.getNotExistedTables()
	if err != nil {
		return fmt.Errorf("trying to get all migrations: %w", err)
	}

	var queue []*queueMigration
	for _, model := range notExisted {
		if m, ok := model.(models.MigrationInterface); ok {
			tc := m.TableCreate()
			tableName := db.getType(m)
			if len(tc.Dependencies) != 0 {
				if db.checkDependencies(notExisted, tc.Dependencies) {
					queue = append(queue, &queueMigration{
						TableCreation: tc,
						TableName:     tableName,
						Model:         m,
					})
					continue
				}
			}

			if err = db.createTable(tc, tableName, m); err != nil {
				log.Printf("%v", err)
			}
		}
	}

	for _, q := range queue {
		if err = db.createTable(q.TableCreation, q.TableName, q.Model); err != nil {
			log.Printf("%v", err)
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

func (db *DbConnection) checkDependencies(notExisted []models.MigrationInterface, dependencies []string) bool {
	for _, ne := range notExisted {
		for _, dep := range dependencies {
			if db.getType(ne) == dep {
				return true
			}
		}
	}
	return false
}

func (db *DbConnection) createTable(tc *models.TableCreation, tableName string, m models.MigrationInterface) error {
	_, err := db.Connection.Exec(tc.Sql)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Printf("[DBSTORAGE] migrations for: %s ALREADY EXISTS", tableName)
			return nil
		}
		return fmt.Errorf("cannot make migration: %s, %v", tc.Sql, err)
	}

	migration := models.CreateMigrationModel(tableName, tc.Sql)
	mq := migration.CreateRecord()
	_, err = db.Connection.Exec(mq)
	if err != nil {
		return fmt.Errorf("cannot create record in database %s: %w", mq, err)
	}

	log.Printf("[DBSTORAGE] Migration for: %s SUCCESSFUL", db.getType(m))
	return nil
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
