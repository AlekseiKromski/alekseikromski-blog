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

	// GetTags - will return a list of tags
	GetTags() []*models.Tag

	// UpdatePost - will update singe post
	UpdatePost(post *models.Post) error

	// DeletePost - will delete singe post
	DeletePost(id int) error

	// CreateComment - will create a comment for post
	CreateComment(comment *models.Comment) error

	// GetComments - will return comments for single post
	GetComments(id int) []*models.Comment

	// DeleteCategory - will delete singe category
	DeleteCategory(id int) error

	// UpdateCategory - will update singe category
	UpdateCategory(category *models.Category) error

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
