package router

import (
	"fmt"
	"net/http"
	"strings"
)

type RouteEntityCreation interface {
	CreateEntity(path, method string, handler http.HandlerFunc)
}

type Group struct {
	router *Router
	prefix string
}

func (g *Group) CreateEntity(path, method string, handler http.HandlerFunc) {
	path = strings.TrimLeft(path, "/")
	entity := RouterEntity{
		fmt.Sprintf("%s%s", g.prefix, path),
		method,
		handler,
	}

	g.router.registerRoute(&entity)
}

// Router - implementation of router with dynamic urls
type Router struct {
	registeredRoutes []*RouterEntity
}

func NewRouter() *Router {
	return &Router{}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, e := range router.registeredRoutes {
		match := e.Match(r)
		if !match {
			continue
		}

		e.Handler.ServeHTTP(w, r)
		return
	}
	http.NotFound(w, r)
}

// RegisterRoute - for register already created route entity
func (router *Router) registerRoute(re *RouterEntity) {
	router.registeredRoutes = append(router.registeredRoutes, re)
}

func (router *Router) CreateEntity(path, method string, handler http.HandlerFunc) {
	entity := &RouterEntity{
		Path:    path,
		Method:  method,
		Handler: handler,
	}

	router.registerRoute(entity)
}

func (r *Router) CreateGroup(prefix string) *Group {
	if prefix[len(prefix)-1:] != "/" {
		prefix += "/"
	}

	return &Group{router: r, prefix: prefix}
}

// RouterEntity - one model for creating routes
type RouterEntity struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Match - check the match for entity
func (re *RouterEntity) Match(r *http.Request) bool {
	if r.Method != re.Method {
		return false
	}

	if r.URL.Path != re.Path {
		return false
	}

	return true
}
