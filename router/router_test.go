package router

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterRegistration(t *testing.T) {
	router := NewRouter()

	router.CreateEntity("/healthz", http.MethodGet, func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("200 OK"))
	})

	server := httptest.NewServer(router)
	defer server.Close()

	response, err := http.DefaultClient.Get(fmt.Sprintf("%s/healthz", server.URL))
	assert.NoError(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "200 OK", string(responseBytes))
}

func TestTest(t *testing.T) {
	router := NewRouter()

	v1 := router.CreateGroup("/v1")
	{
		v1.CreateEntity("/healthz", http.MethodGet, func(writer http.ResponseWriter, request *http.Request) { writer.Write([]byte("200 OK")) })
	}

	server := httptest.NewServer(router)
	defer server.Close()

	response, err := http.DefaultClient.Get(fmt.Sprintf("%s/v1/healthz", server.URL))
	assert.NoError(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "200 OK", string(responseBytes))
}
