package e2e

import (
	"alekseikromski.com/blog/api"
	"alekseikromski.com/blog/api/storage/dbstore"
	v1 "alekseikromski.com/blog/api/v1"
	"alekseikromski.com/blog/router"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

// TestShouldReturnRecords - check, that method return right order of posts (order by createdAt)
func TestShouldReturnRecords(t *testing.T) {
	router := router.NewRouter()
	dbstore, err := dbstore.NewDbConnection(
		"postgres",
		"postgres",
		"localhost",
		"5432",
		"blog",
	)
	assert.NoError(t, err)

	for _, api := range []api.Api{
		v1.NewV1(dbstore, router),
	} {
		api.RegisterRoutes()
	}

	server := httptest.NewServer(router)
	defer server.Close()

	actual := DoRequest(t, fmt.Sprintf("%s/v1/get-last-posts/5/0", server.URL))
	expected := getFile("posts_expected_1.json")
	assert.JSONEq(t, expected, actual)

	actual = DoRequest(t, fmt.Sprintf("%s/v1/get-last-posts/5/5", server.URL))
	expected = getFile("posts_expected_2.json")
	assert.JSONEq(t, expected, actual)
}

func DoRequest(t *testing.T, url string) string {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("cannot do request: %v", err)
	}
	assert.Equal(t, 200, response.StatusCode)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("cannot read response body: %v", err)
	}

	log.Printf("received content: %s", string(content))

	return string(content)
}

func getFile(filename string) string {
	file, err := os.Open(path.Join("testdata", filename))
	if err != nil {
		log.Fatalf("cannot get file: %v", err)
	}
	defer file.Close()

	filecontent, err := io.ReadAll(file)
	return string(filecontent)
}
