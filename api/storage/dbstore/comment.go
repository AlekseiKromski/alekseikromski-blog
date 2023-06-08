package dbstore

import (
	"alekseikromski.com/blog/api/storage/models"
	"fmt"
	"log"
)

func (db *DbConnection) CreateComment(comment *models.Comment) (*models.Comment, error) {
	comment = models.CreateCommentWithData(comment.Name, comment.Text, comment.PostID)

	query := comment.CreateRecord()
	result, err := db.Connection.Exec(query)
	if err != nil {
		log.Printf("Problem: %v", err)
		return nil, fmt.Errorf("cannot create comment record: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("cannot get info about creation: %w", err)
	}

	if affected == 0 {
		return nil, fmt.Errorf("new comment was not created. Affected rows: %d, sql: %s", affected, query)
	}

	return comment, nil
}

func (db *DbConnection) GetComments(post_id int) []*models.Comment {
	comments := []*models.Comment{}

	query := models.GetComments(post_id)

	log.Printf("[DBSTORE] running query: %s", query)

	rows, err := db.Connection.Query(query)
	if err != nil {
		return comments
	}

	for rows.Next() {
		comment := models.CreateComment()
		var scanError error
		scanError = rows.Scan(
			&comment.ID,
			&comment.Name,
			&comment.Text,
			&comment.PostID,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.DeletedAt,
		)

		if scanError != nil {
			log.Printf("troubles during scanning: %w", err)
		}

		comments = append(comments, comment)
	}

	return comments
}
