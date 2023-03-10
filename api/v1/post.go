package v1

import "net/http"

func (v *v1) GetLastPosts(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("200 OK"))
}
