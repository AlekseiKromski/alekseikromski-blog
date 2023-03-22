package router

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Params map[string]string

// Router - implementation of router with dynamic urls
type Router struct {
	registeredRoutes []*Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	url := strings.Trim(request.URL.Path, "/")
	pathList := strings.Split(url, "/")

	for _, e := range r.registeredRoutes {
		if !e.Match(pathList, e.Method) {
			continue
		}

		params := Params{
			"test": "hello from params",
		}

		ctx := context.WithValue(request.Context(), "params", params)
		test := request.WithContext(ctx)

		e.Handler.ServeHTTP(w, test)
		return
	}
	http.NotFound(w, request)
}

// CreateRoute - create new route with all parameters
func (r *Router) CreateRoute(path, method string, handler http.HandlerFunc) {
	entity := &Route{
		Path:     path,
		Method:   method,
		Handler:  handler,
		PathList: nil,
	}

	r.registerRoute(entity)
}

// CreateGroup - create a group of router with the same prefix
func (r *Router) CreateGroup(prefix string) *Group {
	if prefix[len(prefix)-1:] != "/" {
		prefix += "/"
	}

	return &Group{router: r, prefix: prefix}
}

// RegisterRoute - for register already created route entity
func (r *Router) registerRoute(re *Route) {
	path := strings.Trim(re.Path, "/")
	pathList := strings.Split(path, "/")

	dynamicPathRegexp, err := regexp.Compile(`{\w+}+`)
	if err != nil {
		log.Fatalf("Cannot create dynamic regexp: %v", err)
	}

	newPathList := make([]*regexp.Regexp, len(pathList))
	for i, item := range pathList {

		var newItem *regexp.Regexp

		location := dynamicPathRegexp.FindStringIndex(item)
		if len(location) != 0 {
			newItem = regexp.MustCompile(`\S+`)
			newPathList[i] = newItem
			continue
		}

		newItem = regexp.MustCompile(item)
		newPathList[i] = newItem
	}

	re.PathList = newPathList

	r.registeredRoutes = append(r.registeredRoutes, re)
}
