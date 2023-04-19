package v1

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/router"
	"encoding/json"
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
		"/get-last-posts/{size}/{indent}",
		http.MethodGet,
		v.GetLastPosts,
	)
	group.CreateRoute(
		"/post/get-last-posts-by-category/{category_id}/{size}/{indent}",
		http.MethodGet,
		v.GetLastPostsByCategory,
	)
	group.CreateRoute(
		"/post/get-post/{id}",
		http.MethodGet,
		v.GetSinglePost,
	)
	group.CreateRoute(
		"/create-post",
		http.MethodPost,
		v.CreatePost,
	)
}

func (v *v1) ReturnErrorResponse(err error, w http.ResponseWriter) {
	w.WriteHeader(ClassifyError(err))
	json.NewEncoder(w).Encode(err)
}

func (v *v1) ReturnResponse(w http.ResponseWriter, payload []byte) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if payload != nil {
		w.Write(payload)
	}
}
