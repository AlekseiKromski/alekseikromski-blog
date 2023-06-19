package router

import (
	"alekseikromski.com/blog/api/guard"
	"context"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type Params map[string]string

// Router - implementation of router with dynamic urls
type Router struct {
	registeredRoutes []*Route
	guards           map[string]guard.Guard
}

func NewRouter(guards []guard.Guard) *Router {
	return &Router{
		guards: guardParsing(guards),
	}
}

func guardParsing(guards []guard.Guard) map[string]guard.Guard {
	gs := make(map[string]guard.Guard, len(guards))
	for _, guard := range guards {
		gs[reflect.TypeOf(guard).Elem().Name()] = guard
	}

	return gs
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	//cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if request.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	url := strings.Trim(request.URL.Path, "/")
	pathList := strings.Split(url, "/")

	for _, e := range r.registeredRoutes {
		match, params := e.Match(pathList, e.Method)
		if !match {
			continue
		}

		ctx := context.WithValue(request.Context(), "params", params)
		request = request.WithContext(ctx)

		//Guard
		if candidate := r.guards[e.Guard]; candidate != nil {
			if !candidate.Check(request) {
				log.Printf("[INFO] Guard check failed for: %s", e.Path)
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}

		//Middlewares
		if len(e.Middlewares) != 0 {
			for _, middleware := range e.Middlewares {
				if passed := middleware(request, w); !passed {
					return
				}
			}
		}

		e.Handler.ServeHTTP(w, request)
		return
	}

	//By default, return index.html
	file, _ := os.ReadFile("./front-end/build/index.html")
	w.Write(file)
}

// CreateRoute - create new route with all parameters
func (r *Router) CreateRoute(path, method string, handler http.HandlerFunc, guard *string, middlewares ...Middleware) {
	g := ""
	if guard != nil {
		g = *guard
	}
	entity := &Route{
		Path:        path,
		Method:      method,
		Guard:       g,
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

func (r *Router) GetParams(request *http.Request) Params {
	return request.Context().Value("params").(Params)
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
	log.Printf("[ROTUER] new route was registers: %s [%s][GUARD:%s]", re.Path, re.Method, re.Guard)
}
