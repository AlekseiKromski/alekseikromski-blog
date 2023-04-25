package v1

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/router"
	"encoding/json"
	"net/http"
)

type V1 struct {
	Version string
	router  *router.Router
	storage storage.Storage
}

func NewV1(storage storage.Storage, router *router.Router) *V1 {
	return &V1{
		Version: "V1",
		router:  router,
		storage: storage,
	}
}

func (v *V1) RegisterRoutes() {
	group := v.router.CreateGroup("/V1/")
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
		"/category/all",
		http.MethodGet,
		v.GetAllCategories,
	)
	group.CreateRoute(
		"/create-post",
		http.MethodPost,
		v.CreatePost,
	)
}

func (v *V1) ReturnErrorResponse(err error, w http.ResponseWriter) {
	w.WriteHeader(ClassifyError(err))
	json.NewEncoder(w).Encode(err)
}

func (v *V1) ReturnResponse(w http.ResponseWriter, payload []byte) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if payload != nil {
		w.Write(payload)
	}
}
