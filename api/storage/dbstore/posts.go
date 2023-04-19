package dbstore

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/storage/models"
	"fmt"
	"log"
)

func (db *DbConnection) GetPosts(request *storage.QueryRequest) []*models.Post {
	var posts []*models.Post

	if request.ID != nil {
		query := models.GetPost(*request.ID)

		log.Printf("[DBSTORE] running query: %s", query)

		rows, err := db.Connection.Query(query)
		if err != nil {
			return posts
		}

		for rows.Next() {
			post := models.CreatePost()
			var scanError error
			category := models.CreateCategory()
			scanError = rows.Scan(
				&post.ID,
				&post.Title,
				&post.Description,
				&post.CategoryID,
				&post.CreatedAt,
				&post.UpdatedAt,
				&post.DeletedAt,
				&category.ID,
				&category.Name,
				&category.CreatedAt,
				&category.UpdatedAt,
				&category.DeletedAt,
			)

			post.Category = category

			if scanError != nil {
				log.Printf("troubles during scanning: %w", err)
			}

			posts = append(posts, post)
		}

		return posts
	}

	if request.Limit > 0 {
		query, withCategory := models.GetLastPosts(request.Limit, request.Offset, request.CategoryID)

		log.Printf("[DBSTORE] running query: %s", query)

		rows, err := db.Connection.Query(query)
		if err != nil {
			return posts
		}

		for rows.Next() {
			post := models.CreatePost()
			var scanError error
			if withCategory {
				category := models.CreateCategory()
				scanError = rows.Scan(
					&post.ID,
					&post.Title,
					&post.Description,
					&post.CategoryID,
					&post.CreatedAt,
					&post.UpdatedAt,
					&post.DeletedAt,
					&category.ID,
					&category.Name,
					&category.CreatedAt,
					&category.UpdatedAt,
					&category.DeletedAt,
				)
				post.Category = category
			} else {
				scanError = rows.Scan(
					&post.ID,
					&post.Title,
					&post.Description,
					&post.CategoryID,
					&post.CreatedAt,
					&post.UpdatedAt,
					&post.DeletedAt,
				)
			}

			if scanError != nil {
				log.Printf("troubles during scanning: %w", err)
			}

			posts = append(posts, post)
		}
	}

	return posts
}

func (db *DbConnection) CreatePost(post *models.Post) (bool, error) {

	//Recreate from json model
	post = models.CreatePostWithData(post.Title, post.Description)

	query := post.CreateRecord()

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
