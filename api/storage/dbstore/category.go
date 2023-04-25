package dbstore

import (
	"alekseikromski.com/blog/api/storage/models"
	"log"
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
