package router

import (
	"fmt"
	"net/http"
	"strings"
)

type Group struct {
	router *Router
	prefix string
}

func (g *Group) CreateRoute(path, method string, handler http.HandlerFunc, middlewares ...Middleware) {
	path = strings.TrimLeft(path, "/")
	entity := Route{
		fmt.Sprintf("%s%s", g.prefix, path),
		method,
		false,
		handler,
		nil,
		middlewares,
	}

	g.router.registerRoute(&entity)
}
