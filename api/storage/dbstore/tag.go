package dbstore

import (
	"alekseikromski.com/blog/api/storage/models"
	"log"
)

func (db *DbConnection) GetTags(postID *int) []*models.Tag {
	var tags []*models.Tag

	query := models.GetTags(postID)

	log.Printf("[DBSTORE] running query: %s", query)

	rows, err := db.Connection.Query(query)
	if err != nil {
		log.Printf("Problem: %v", err)
		return tags
	}

	for rows.Next() {
		var scanError error
		tag := models.CreateTag()

		scanError = rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.PostID,
			&tag.CreatedAt,
			&tag.UpdatedAt,
			&tag.DeletedAt,
		)

		if scanError != nil {
			log.Printf("troubles during scanning: %v", err)
		}

		tags = append(tags, tag)
	}

	return tags
}
