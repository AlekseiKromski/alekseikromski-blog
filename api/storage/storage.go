package storage

import (
	"alekseikromski.com/blog/api/storage/models"
)

type Migration interface {
	RunMigrations() error
}

type Storage interface {
	// GetPosts - will return all posts by query request
	GetPosts(request *QueryRequest) []*models.Post

	// CreatePost - will create a post
	CreatePost(post *models.Post) (bool, error)

	// GetCategories - will return a list of categories
	GetCategories() []*models.Category

	//General functions
	Stop()
}

type QueryRequest struct {
	ID         *int
	Limit      int
	Offset     int
	CategoryID int
}

// NewQueryRequest - will return basic object without any properties
func NewQueryRequest() *QueryRequest {
	return &QueryRequest{}
}
