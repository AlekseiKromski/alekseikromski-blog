package v1

import (
	"fmt"
	"log"
	"net/http"
)

type v1 struct {
	Version string
	routes  map[string]http.HandlerFunc
}

func NewV1() *v1 {
	return &v1{
		Version: "v1",
	}
}

func (v *v1) RegisterRoutes() {
	v.routes = map[string]http.HandlerFunc{
		fmt.Sprintf("/%s/get-last-posts", v.Version): v.GetLastPosts,
	}
}

func (v *v1) Mount(mux *http.ServeMux) {
	for route, handler := range v.routes {
		log.Printf("Route [ %s ] was mounted - V1", route)
		mux.HandleFunc(route, handler)
	}
}
