package dbstore

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/storage/models"
	"fmt"
	"log"
	"time"
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
				&post.Img,
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

	if request.Limit > 0 {
		query := models.GetLastPosts(request.Limit, request.Offset, request.CategoryID)

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
				&post.Img,
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
	}

	return posts
}

func (db *DbConnection) UpdatePost(post *models.Post) error {
	post.UpdatedAt = time.Now().Format(time.RFC3339)
	query := models.UpdatePost(post)
	log.Printf("[DBSTORE] running query: %s", query)
	_, err := db.Connection.Query(query)
	if err != nil {
		return fmt.Errorf("cannot update post: %w", err)
	}
	return nil
}

func (db *DbConnection) CreatePost(post *models.Post) (bool, error) {

	//Recreate from json model
	post = models.CreatePostWithData(post.Title, post.Description, "", post.CategoryID)

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

func (db *DbConnection) DeletePost(id int) error {
	request := storage.NewQueryRequest()
	request.ID = &id
	posts := db.GetPosts(request)
	if len(posts) != 0 && len(posts) != 1 {
		return fmt.Errorf("cannot get post by id")
	}

	post := posts[0]
	post.Soft()

	if err := db.UpdatePost(post); err != nil {
		fmt.Errorf("cannot update object: %w", err)
	}
	return nil
}
