package v1

import (
	"alekseikromski.com/blog/api/storage"
	"fmt"
	"log"
	"net/http"
)

type v1 struct {
	Version string
	routes  map[string]http.HandlerFunc
	storage storage.Storage
}

func NewV1(storage storage.Storage) *v1 {
	return &v1{
		Version: "v1",
		storage: storage,
	}
}

func (v *v1) RegisterRoutes() {
	v.routes = map[string]http.HandlerFunc{
		fmt.Sprintf("/%s/get-last-posts", v.Version): v.GetLastPosts,
		fmt.Sprintf("/%s/create-post", v.Version):    v.CreatePost,
	}
}

func (v *v1) Mount(mux *http.ServeMux) {
	for route, handler := range v.routes {
		log.Printf("Route [ %s ] was mounted - V1", route)
		mux.Handle(route, handler)
	}
}
