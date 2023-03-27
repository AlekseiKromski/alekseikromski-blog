package router

import "net/http"

type Middleware func(r *http.Request, w http.ResponseWriter) bool
