package v1

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/router"
	"net/http"
)

type v1 struct {
	Version string
	router  *router.Router
	storage storage.Storage
}

func NewV1(storage storage.Storage, router *router.Router) *v1 {
	return &v1{
		Version: "v1",
		router:  router,
		storage: storage,
	}
}

func (v *v1) RegisterRoutes() {
	group := v.router.CreateGroup("/v1/")
	group.CreateRoute(
		"/get-last-posts",
		http.MethodGet,
		v.GetLastPosts,
	)
	group.CreateRoute(
		"/create-post",
		http.MethodPost,
		v.CreatePost,
	)
}
