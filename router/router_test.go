package router

import (
	"alekseikromski.com/blog/api/guard"
	"alekseikromski.com/blog/api/guard/jwt"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter([]guard.Guard{})

	router.CreateRoute("/healthz", http.MethodGet, func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("200 OK"))
	}, nil)

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

func TestGroup_CreateRouter(t *testing.T) {
	router := NewRouter([]guard.Guard{})

	v1 := router.CreateGroup("/v1")
	{
		v1.CreateRoute("/healthz", http.MethodGet, func(writer http.ResponseWriter, request *http.Request) { writer.Write([]byte("200 OK")) }, nil)
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

func TestRouter_CreateDynamicRoute(t *testing.T) {
	router := NewRouter([]guard.Guard{})

	router.CreateRoute("/v1/{id}/{test}/test", http.MethodGet, func(writer http.ResponseWriter, request *http.Request) {
		params := router.GetParams(request)
		assert.Equal(t, "f8aef97f-60aa-42de-b7b1-db8f5d45f6fd", params["id"])
		writer.Write([]byte("200 OK"))
	}, nil)

	server := httptest.NewServer(router)
	defer server.Close()

	response, err := http.DefaultClient.Get(fmt.Sprintf("%s/v1/f8aef97f-60aa-42de-b7b1-db8f5d45f6fd/2b833c3d-289b-4783-b0f9-313e44eb11e7/test", server.URL))
	assert.NoError(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "200 OK", string(responseBytes))
}

func TestRouter_CreateDynamicRouteWithMiddleware(t *testing.T) {
	router := NewRouter([]guard.Guard{})

	router.CreateRoute(
		"/v1/{id}/{test}/test",
		http.MethodGet,
		func(writer http.ResponseWriter, request *http.Request) { writer.Write([]byte("200 OK")) },
		nil,
		func(r *http.Request, w http.ResponseWriter) bool {
			//Middleware logic
			params := router.GetParams(r)
			if "f8aef97f-60aa-42de-b7b1-db8f5d45f6fd" == params["id"] {
				w.WriteHeader(http.StatusBadRequest)
				return false
			}

			return true
		},
	)

	server := httptest.NewServer(router)
	defer server.Close()

	response, err := http.DefaultClient.Get(fmt.Sprintf("%s/v1/f8aef97f-60aa-42de-b7b1-db8f5d45f6fd/2b833c3d-289b-4783-b0f9-313e44eb11e7/test", server.URL))
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestRouter_CreateGroupRouteWithMiddleware(t *testing.T) {
	router := NewRouter([]guard.Guard{})

	group := router.CreateGroup("/v1/")
	group.CreateRoute(
		"/{id}/{test}/test",
		http.MethodGet,
		func(writer http.ResponseWriter, request *http.Request) { writer.Write([]byte("200 OK")) },
		nil,
		func(r *http.Request, w http.ResponseWriter) bool {
			//Middleware logic
			params := router.GetParams(r)
			if "f8aef97f-60aa-42de-b7b1-db8f5d45f6fd" == params["id"] {
				w.WriteHeader(http.StatusBadRequest)
				return false
			}

			return true
		},
	)

	server := httptest.NewServer(router)
	defer server.Close()

	response, err := http.DefaultClient.Get(fmt.Sprintf("%s/v1/f8aef97f-60aa-42de-b7b1-db8f5d45f6fd/2b833c3d-289b-4783-b0f9-313e44eb11e7/test", server.URL))
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestRouter_CreateGroupRouteWithGuard(t *testing.T) {
	jwtGuard := jwt.New("secret")
	guards := []guard.Guard{
		jwtGuard,
	}
	router := NewRouter(guards)

	guard := "JWTGuard"
	group := router.CreateGroup("/v1/")
	group.CreateRoute(
		"/{id}/{test}/test",
		http.MethodGet,
		func(writer http.ResponseWriter, request *http.Request) { writer.Write([]byte("200 OK")) },
		&guard,
	)

	server := httptest.NewServer(router)
	defer server.Close()

	token := jwtGuard.Auth(1)
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/f8aef97f-60aa-42de-b7b1-db8f5d45f6fd/2b833c3d-289b-4783-b0f9-313e44eb11e7/test", server.URL), nil)
	assert.NoError(t, err)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestRouter_CreateGroupRouteWithGuard_shouldFail(t *testing.T) {
	guards := []guard.Guard{
		jwt.New("secret"),
	}
	router := NewRouter(guards)

	guard := "JWTGuard"
	group := router.CreateGroup("/v1/")
	group.CreateRoute(
		"/{id}/{test}/test",
		http.MethodGet,
		func(writer http.ResponseWriter, request *http.Request) { writer.Write([]byte("200 OK")) },
		&guard,
	)

	server := httptest.NewServer(router)
	defer server.Close()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/f8aef97f-60aa-42de-b7b1-db8f5d45f6fd/2b833c3d-289b-4783-b0f9-313e44eb11e7/test", server.URL), nil)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "test-token"))

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, response.StatusCode)
}
