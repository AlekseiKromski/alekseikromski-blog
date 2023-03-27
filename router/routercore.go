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
		match, params := e.Match(pathList, e.Method)
		if !match {
			continue
		}

		ctx := context.WithValue(request.Context(), "params", params)
		request = request.WithContext(ctx)

		if len(e.Middlewares) != 0 {
			for _, middleware := range e.Middlewares {
				passed := middleware(request, w)
				if !passed {
					return
				}
			}
		}

		e.Handler.ServeHTTP(w, request)
		return
	}
	http.NotFound(w, request)
}

// CreateRoute - create new route with all parameters
func (r *Router) CreateRoute(path, method string, handler http.HandlerFunc, middlewares ...Middleware) {
	entity := &Route{
		Path:        path,
		Method:      method,
		IsAll:       false,
		Handler:     handler,
		PathList:    nil,
		Middlewares: middlewares,
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

	newPathList := make([]*Param, len(pathList))
	for i, item := range pathList {
		param := &Param{
			Key: item,
		}

		location := dynamicPathRegexp.FindStringIndex(item)
		if len(location) != 0 {
			param.Key = re.prepareDynamicItemKey(item)
			param.RegExp = regexp.MustCompile(`\S+`)
			param.IsDynamic = true
			newPathList[i] = param
			continue
		}

		if item == "*" {
			param.RegExp = regexp.MustCompile(`.+`)
			param.IsDynamic = false
			newPathList[i] = param
			break
		}

		if item == "" {
			re.IsAll = true
		}

		param.RegExp = regexp.MustCompile(item)
		newPathList[i] = param
	}

	re.PathList = newPathList

	r.registeredRoutes = append(r.registeredRoutes, re)
	log.Printf("[ROTUER] new route was registers: %s", re.Path)
}

func (r *Router) GetParams(request *http.Request) Params {
	return request.Context().Value("params").(Params)
}
