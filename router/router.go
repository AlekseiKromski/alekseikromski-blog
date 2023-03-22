package router

import (
	"net/http"
)

type RouteEntityCreation interface {
	CreateRoute(path, method string, handler http.HandlerFunc)
}
