package api

type Api interface {
	RegisterRoutes() error // RegisterRoutes - need to register all routes, that we have in api realization
}
