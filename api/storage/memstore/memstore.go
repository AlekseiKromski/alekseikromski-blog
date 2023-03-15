package memstore

import (
	"alekseikromski.com/blog/api/storage"
	"fmt"
	"log"
)

type Memstorage struct {
	posts []*storage.Post
}

func NewMemStorage() *Memstorage {
	return &Memstorage{}
}

func (m *Memstorage) GetPosts(request *storage.QueryRequest) []*storage.Post {
	var posts []*storage.Post
	if request.Limit != 0 {
		var buffer []*storage.Post
		if len(m.posts) > request.Limit {
			for i := len(m.posts) - 1; len(buffer) < request.Limit; i-- {
				buffer = append(buffer, m.posts[i])
			}
		} else {
			buffer = m.posts
		}

		posts = buffer
	}
	log.Println("Was sent last posts")
	return posts
}

func (m *Memstorage) CreatePost(post *storage.Post) (bool, error) {
	if !post.Validate() {
		return false, fmt.Errorf("post is not valid")
	}

	m.posts = append(m.posts, post)
	return true, nil
}
