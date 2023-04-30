package dbstore

import (
	"alekseikromski.com/blog/api/storage/models"
	"fmt"
	"log"
	"time"
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

func (db *DbConnection) UpdateTag(tag *models.Tag) error {
	tag.UpdatedAt = time.Now().Format(time.RFC3339)
	query := models.UpdateTag(tag)
	log.Printf("[DBSTORE] running query: %s", query)
	_, err := db.Connection.Query(query)
	if err != nil {
		return fmt.Errorf("cannot update post: %w", err)
	}
	return nil
}

func (db *DbConnection) DeleteTag(id int) error {
	tag := models.CreateTag()

	query := models.GetTagByID(id)

	log.Printf("[DBSTORE] running query: %s", query)

	rows, err := db.Connection.Query(query)
	if err != nil {
		return fmt.Errorf("cannot find category with id %d, %w", id, err)
	}

	for rows.Next() {
		var scanError error
		scanError = rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.PostID,
			&tag.CreatedAt,
			&tag.UpdatedAt,
			&tag.DeletedAt,
		)

		if scanError != nil {
			return fmt.Errorf("troubles during scanning: %w", err)
		}
	}

	tag.Soft()

	if err := db.UpdateTag(tag); err != nil {
		fmt.Errorf("cannot update object: %w", err)
	}
	return nil
}
