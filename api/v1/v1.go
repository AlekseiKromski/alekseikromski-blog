package v1

import (
	"alekseikromski.com/blog/api/guard"
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/router"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
)

type V1 struct {
	Version string
	router  *router.Router
	storage storage.Storage
	guards  map[string]guard.Guard
}

func NewV1(storage storage.Storage, router *router.Router, gs []guard.Guard) *V1 {
	return &V1{
		Version: "V1",
		router:  router,
		storage: storage,
		guards:  parseGuards(gs),
	}
}

func parseGuards(guards []guard.Guard) map[string]guard.Guard {
	gs := make(map[string]guard.Guard, len(guards))
	for _, guard := range guards {
		gs[reflect.TypeOf(guard).Elem().Name()] = guard
	}

	return gs
}

func (v *V1) RegisterRoutes() error {
	fp, err := filepath.Abs(filepath.Join("api", "v1", "data", "routes.json"))
	if err != nil {
		return fmt.Errorf("cannot find routes file: %w", err)
	}
	p := NewParser(fp, reflect.ValueOf(v))
	if err := p.Parse(v.router); err != nil {
		return fmt.Errorf("cannot parse routes file: %w", err)
	}
	return nil
}

func (v *V1) ReturnErrorResponse(err error, w http.ResponseWriter) {
	w.WriteHeader(ClassifyError(err))
	json.NewEncoder(w).Encode(err)
}

func (v *V1) ReturnResponse(w http.ResponseWriter, payload []byte) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if payload != nil {
		w.Write(payload)
	}
}
