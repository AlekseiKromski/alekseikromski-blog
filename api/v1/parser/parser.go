package parser

import (
	"alekseikromski.com/blog/router"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
)

type Parsing interface {
	// Parse - will return extended router with configure routes from file (.json file)
	Parse(router *router.Router) error
}

type route struct {
	Method  string `json:"method"`
	Route   string `json:"route"`
	Handler string `json:"handler"`
}

type group struct {
	Url    string   `json:"url"`
	Routes []*route `json:"routes"`
}

type Model struct {
	Groups []*group `json:"groups"`
	Routes []*route `json:"routes"`
}

type Parser struct {
	// fp - filepath
	fp  string
	api reflect.Value
}

func NewParser(fp string, api reflect.Value) *Parser {
	return &Parser{
		fp:  fp,
		api: api,
	}
}

func (p *Parser) Parse(router *router.Router) error {
	f, err := os.Open(p.fp)
	if err != nil {
		return fmt.Errorf("cannot open rotues file: %w", err)
	}
	defer f.Close()

	var model *Model
	err = json.NewDecoder(f).Decode(&model)
	if err != nil {
		return fmt.Errorf("cannot decode rotues file: %w", err)
	}

	//Default Routes (not group)
	for _, route := range model.Routes {
		p.registerRouter(router, route, p.api.MethodByName(route.Handler))
	}

	//Group routes
	for _, group := range model.Groups {
		g := router.CreateGroup(group.Url)
		for _, route := range group.Routes {
			p.registerRouter(g, route, p.api.MethodByName(route.Handler))
		}
	}

	return nil
}

func (p *Parser) registerRouter(re router.RouteEntityCreation, route *route, handler reflect.Value) {
	re.CreateRoute(route.Route, route.Method, func(writer http.ResponseWriter, request *http.Request) {
		handler.Call([]reflect.Value{reflect.ValueOf(writer), reflect.ValueOf(request)})
	})
}
