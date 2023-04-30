package dbstore

import (
	"alekseikromski.com/blog/api/storage/models"
	"fmt"
	"log"
	"time"
)

func (db *DbConnection) GetCategories() []*models.Category {
	var categories []*models.Category

	query := models.GetCategories()

	log.Printf("[DBSTORE] running query: %s", query)

	rows, err := db.Connection.Query(query)
	if err != nil {
		return categories
	}

	for rows.Next() {
		var scanError error
		category := models.CreateCategory()

		scanError = rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.DeletedAt,
		)

		if scanError != nil {
			log.Printf("troubles during scanning: %w", err)
		}

		categories = append(categories, category)
	}

	return categories
}

func (db *DbConnection) DeleteCategory(id int) error {
	category := models.CreateCategory()

	query := models.GetCategory(id)

	log.Printf("[DBSTORE] running query: %s", query)

	rows, err := db.Connection.Query(query)
	if err != nil {
		return fmt.Errorf("cannot find category with id %d, %w", id, err)
	}

	for rows.Next() {
		var scanError error
		scanError = rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.DeletedAt,
		)

		if scanError != nil {
			return fmt.Errorf("troubles during scanning: %w", err)
		}
	}

	category.Soft()

	if err := db.UpdateCategory(category); err != nil {
		fmt.Errorf("cannot update object: %w", err)
	}
	return nil
}

func (db *DbConnection) UpdateCategory(category *models.Category) error {
	category.UpdatedAt = time.Now().Format(time.RFC3339)
	query := models.UpdateCategory(category)
	log.Printf("[DBSTORE] running query: %s", query)
	_, err := db.Connection.Query(query)
	if err != nil {
		return fmt.Errorf("cannot update post: %w", err)
	}
	return nil
}

func (db *DbConnection) CreateCategory(category *models.Category) (bool, error) {

	//Recreate from json model
	category = models.CreateCategoryWithData(category.Name)

	query := category.CreateRecord()

	result, err := db.Connection.Exec(query)
	if err != nil {
		return false, fmt.Errorf("cannot create post record: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("cannot get info about creation: %w", err)
	}

	if affected == 0 {
		return false, fmt.Errorf("new was not created. Affected rows: %d, sql: %s", affected, query)
	}

	return true, nil
}
