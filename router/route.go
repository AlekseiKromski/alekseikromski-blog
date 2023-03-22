package router

import (
	"net/http"
	"regexp"
)

// Route - one model for creating routes
type Route struct {
	Path     string
	Method   string
	Handler  http.HandlerFunc
	PathList []*regexp.Regexp
}

// Match - check the match for entity
func (ro *Route) Match(r []string, method string) bool {
	if method != ro.Method {
		return false
	}

	if len(r) == len(ro.PathList) {

		for index, item := range ro.PathList {
			if !item.Match([]byte(r[index])) {
				return false
			}
		}

		return true
	}

	return false
}
