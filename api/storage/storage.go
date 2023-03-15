package storage

type Storage interface {
	// GetPosts - will return all posts
	GetPosts(request *QueryRequest) []*Post
	CreatePost(post *Post) (bool, error)
}

type QueryRequest struct {
	ID    int
	Limit int
}

// NewQueryRequest - will return basic object without any properties
func NewQueryRequest() *QueryRequest {
	return &QueryRequest{}
}
