package api

import "net/http"

type Api interface {
	Mount(mux *http.ServeMux) // Mount - need to mount all handlers (depending on api version)
	RegisterRoutes()          // RegisterRoutes - need to register all routes, that we have in api realization
}
