package v1

import (
	"alekseikromski.com/blog/api/storage"
	"encoding/json"
	"log"
	"net/http"
)

func (v *v1) GetLastPosts(writer http.ResponseWriter, request *http.Request) {
	query := storage.NewQueryRequest()
	query.Limit = 3

	posts := v.storage.GetPosts(query)

	response, err := json.Marshal(posts)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (v *v1) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var post storage.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Decoding error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	created, err := v.storage.CreatePost(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !created {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("created"))
	log.Printf("Post with id [%d] was created", post.ID)
}
