package api

type Api interface {
	RegisterRoutes(routesFile string) error // RegisterRoutes - need to register all routes, that we have in api realization
}
