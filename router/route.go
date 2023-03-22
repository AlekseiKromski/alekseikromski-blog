package router

import (
	"net/http"
	"regexp"
)

type Param struct {
	Key       string
	IsDynamic bool
	RegExp    *regexp.Regexp
}

// Route - one model for creating routes
type Route struct {
	Path     string
	Method   string
	Handler  http.HandlerFunc
	PathList []*Param
}

// Match - check the match for entity
func (ro *Route) Match(r []string, method string) (bool, Params) {
	if method != ro.Method {
		return false, nil
	}

	if len(r) == len(ro.PathList) {

		for index, item := range ro.PathList {
			if !item.RegExp.Match([]byte(r[index])) {
				return false, nil
			}
		}

		return ro.associateValues(r)
	}

	return false, nil
}

func (ro *Route) associateValues(values []string) (bool, Params) {
	params := Params{}
	for i, item := range ro.PathList {
		if !item.IsDynamic {
			continue
		}
		params[item.Key] = values[i]
	}

	return true, params
}

func (ro *Route) prepareDynamicItemKey(item string) string {
	item = item[1:]
	item = item[:len(item)-1]
	return item
}
