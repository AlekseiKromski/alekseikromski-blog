package v1

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/storage/models"
	"encoding/json"
	"log"
	"net/http"
)

// GetLastPosts
//
//	@Summary		List of last posts
//	@Description	Get last posts from storage
//	@Produce		json
//	@Success		200	{array}	storage.Post
//	@Failure		400
//	@Failure		500
//	@Router			/v1/get-last-posts [get]
func (v *v1) GetLastPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := storage.NewQueryRequest()
	query.Limit = 3

	posts := v.storage.GetPosts(query)

	response, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// CreatePost
//
//	@Summary		Create post
//	@Description	Create a post
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/v1/create-post [post]
func (v *v1) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var post models.Post

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
