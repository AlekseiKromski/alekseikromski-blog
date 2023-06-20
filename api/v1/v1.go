package v1

import (
	"alekseikromski.com/blog/api/guard"
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/router"
	"encoding/json"
	"fmt"
	"net/http"
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

func (v *V1) RegisterRoutes(fp string) error {
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
	w.WriteHeader(http.StatusOK)
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

	if payload != nil {
		w.Write(payload)
	}
}
