package v1

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/v1/parser"
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
}

func NewV1(storage storage.Storage, router *router.Router) *V1 {
	return &V1{
		Version: "V1",
		router:  router,
		storage: storage,
	}
}

func (v *V1) RegisterRoutes() error {
	fp, err := filepath.Abs(filepath.Join("api", "v1", "data", "routes.json"))
	if err != nil {
		return fmt.Errorf("cannot find routes file: %w", err)
	}
	p := parser.NewParser(fp, reflect.ValueOf(v))
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
