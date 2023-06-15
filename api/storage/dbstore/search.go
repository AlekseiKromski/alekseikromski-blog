package dbstore

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/storage/models"
	"log"
)

func (db *DbConnection) Search(sr *storage.SearchRequest) *storage.SearchResult {
	result := &storage.SearchResult{
		Generic: []*models.Post{},
	}

	result.Generic = getPostSearch(sr, db)

	return result
}

func getPostSearch(sr *storage.SearchRequest, db *DbConnection) []*models.Post {
	posts := []*models.Post{}

	query := models.SearchPost(sr.Search)
	rows, err := db.Connection.Query(query)
	if err != nil {
		log.Printf("Problem: %v", err)
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

		//Get comments
		post.Comments = db.GetComments(post.ID)

		//Get tags
		post.Tags = db.GetTags(&post.ID)

		posts = append(posts, post)
	}

	return posts
}
