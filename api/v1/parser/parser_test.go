package parser

import (
	"alekseikromski.com/blog/api/storage/dbstore"
	"alekseikromski.com/blog/api/v1"
	"alekseikromski.com/blog/router"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestParsingJsonFile(t *testing.T) {
	store, _ := dbstore.NewDbConnection("postgres", "postgres", "localhost", "5432", "blog")
	r := router.NewRouter()
	a := v1.NewV1(store, r)

	parser := NewParser("testdata/routes.json", reflect.ValueOf(a))
	err := parser.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	server := httptest.NewServer(r)

	req, err := http.NewRequest(http.MethodGet, server.URL+"/test", nil)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("cannot get response: %v", err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("cannot red response body: %v", err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NotEqual(t, 0, len(content))

	req, err = http.NewRequest(http.MethodGet, server.URL+"/group/test", nil)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("cannot get response: %v", err)
	}
	defer res.Body.Close()

	content, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("cannot red response body: %v", err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NotEqual(t, 0, len(content))
}
