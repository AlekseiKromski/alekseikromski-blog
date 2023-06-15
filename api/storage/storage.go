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

	// CreateCategory - will create a category
	CreateCategory(category *models.Category) (bool, error)

	// CreateTag - will create a tag
	CreateTag(tag *models.Tag) (bool, error)

	// GetCategories - will return a list of categories
	GetCategories() []*models.Category

	// GetTags - will return a list of tags
	GetTags(postID *int) []*models.Tag

	// UpdatePost - will update singe post
	UpdatePost(post *models.Post) error

	// DeletePost - will delete singe post
	DeletePost(id int) error

	// CreateComment - will create a comment for post
	CreateComment(comment *models.Comment) (*models.Comment, error)

	// GetComments - will return comments for single post
	GetComments(id int) []*models.Comment

	// DeleteCategory - will delete singe category
	DeleteCategory(id int) error

	// UpdateCategory - will update single category
	UpdateCategory(category *models.Category) error

	// UpdateTag - will update single tag
	UpdateTag(tag *models.Tag) error

	// DeleteTag - will delete tag
	DeleteTag(id int) error

	// Search - will return search result
	Search(sr *SearchRequest) *SearchResult

	// GetUser - will return user from DB
	GetUser(email string) (*models.User, error)

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

type SearchResult struct {
	Generic []*models.Post `json:"posts"`
}

type SearchRequest struct {
	Search string `json:"search"`
}

// NewSearchRequest - will return basic object without any properties
func NewSearchRequest(search string) *SearchRequest {
	return &SearchRequest{
		Search: search,
	}
}
