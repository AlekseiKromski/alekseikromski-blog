package storage

import (
	"alekseikromski.com/blog/api/storage/models"
)

type Migration interface {
	RunMigrations() error
}

type Storage interface {
	// GetPosts - will return all posts
	GetPosts(request *QueryRequest) []*models.Post
	CreatePost(post *models.Post) (bool, error)

	//General functions
	Stop()
}

type QueryRequest struct {
	ID     int
	Limit  int
	Offset int
}

// NewQueryRequest - will return basic object without any properties
func NewQueryRequest() *QueryRequest {
	return &QueryRequest{}
}
