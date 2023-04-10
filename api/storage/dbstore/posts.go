package dbstore

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/storage/models"
	"fmt"
)

func (db *DbConnection) GetPosts(request *storage.QueryRequest) []*models.Post {
	var posts []*models.Post

	if request.Limit > 0 {
		query := models.GetLastPosts(request.Limit)

		rows, err := db.Connection.Query(query)
		if err != nil {
			return posts
		}

		for rows.Next() {
			post := models.CreatePost()
			rows.Scan(&post.ID, &post.Title, &post.Description, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
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
