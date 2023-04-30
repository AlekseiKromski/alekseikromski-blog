package v1

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/storage/models"
	"alekseikromski.com/blog/router"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// GetLastPosts
//
//	@Summary		List of last posts
//	@Description	Get last posts from storage
//	@Produce		json
//	@Success		200	{array}		models.Post
//	@Failure		400	{object}	V1.JsonError	"if we cannot decode or encode payload"
//	@Failure		500	{object}	V1.InputError	"if we have bad payload"
//	@Router			/V1/get-last-posts [get]
func (v *V1) GetLastPosts(w http.ResponseWriter, r *http.Request) {

	// Get params from context
	ctx := r.Context()
	var params router.Params
	if pr, ok := ctx.Value("params").(router.Params); ok {
		params = pr
	}

	query := storage.NewQueryRequest()

	size, indent, err := v.getSizeAndOffset(params)
	if err != nil {
		v.ReturnErrorResponse(NewInputError(), w)
		return
	}

	query.Limit = size
	query.Offset = indent

	posts := v.storage.GetPosts(query)

	response, err := json.Marshal(posts)
	if err != nil {
		v.ReturnErrorResponse(NewDecodingError(), w)
		return
	}

	v.ReturnResponse(w, response)
}

// GetLastPostsByCategory
//
//	@Summary		List of last posts filtered by category
//	@Description	Get last posts from storage filtered by category
//	@Produce		json
//	@Success		200	{array}		models.Post
//	@Failure		400	{object}	V1.JsonError	"if we cannot decode or encode payload"
//	@Failure		500	{object}	V1.InputError	"if we have bad payload"
//	@Router			/V1/post/get-last-posts-by-category/{category_id}/{size}/{offset} [get]
func (v *V1) GetLastPostsByCategory(w http.ResponseWriter, r *http.Request) {

	// Get params from context
	ctx := r.Context()
	var params router.Params
	if pr, ok := ctx.Value("params").(router.Params); ok {
		params = pr
	}

	query := storage.NewQueryRequest()

	size, indent, err := v.getSizeAndOffset(params)
	if err != nil {
		v.ReturnErrorResponse(NewInputError(), w)
		return
	}

	categoryID, err := strconv.Atoi(params["category_id"])
	if err != nil {
		// Recreate error
		v.ReturnErrorResponse(NewInputError(), w)
		return
	}

	query.Limit = size
	query.Offset = indent
	query.CategoryID = categoryID

	posts := v.storage.GetPosts(query)

	response, err := json.Marshal(posts)
	if err != nil {
		v.ReturnErrorResponse(NewDecodingError(), w)
		return
	}

	v.ReturnResponse(w, response)
}

// GetSinglePost
//
//	@Summary		Return only one post
//	@Description	Get post by id
//	@Produce		json
//	@Success		200	{array}		models.Post
//	@Failure		400	{object}	V1.JsonError	"if we cannot decode or encode payload"
//	@Failure		500	{object}	V1.InputError	"if we have bad payload"
//	@Router			/V1/post/get-post/1 [get]
func (v *V1) GetSinglePost(w http.ResponseWriter, r *http.Request) {

	// Get params from context
	ctx := r.Context()
	var params router.Params
	if pr, ok := ctx.Value("params").(router.Params); ok {
		params = pr
	}

	query := storage.NewQueryRequest()

	postID, err := strconv.Atoi(params["id"])
	if err != nil {
		// Recreate error
		v.ReturnErrorResponse(NewInputError(), w)
		return
	}

	query.ID = &postID

	posts := v.storage.GetPosts(query)

	if len(posts) == 0 {
		v.ReturnErrorResponse(NewInputError(), w)
		return
	}

	//return only one post
	response, err := json.Marshal(posts[0])
	if err != nil {
		v.ReturnErrorResponse(NewDecodingError(), w)
		return
	}

	v.ReturnResponse(w, response)
}

// UpdatePost
//
//	@Summary		Update post
//	@Description	Update single post
//	@Produce		json
//	@Success		200	{array}		models.Post
//	@Failure		400	{object}	V1.JsonError	"if we cannot decode or encode payload"
//	@Failure		500	{object}	V1.InputError	"if we have bad payload"
//	@Router			/v1/post/edit-post [post]
func (v *V1) UpdatePost(w http.ResponseWriter, r *http.Request) {
	postForUpdate := models.CreatePost()

	err := json.NewDecoder(r.Body).Decode(&postForUpdate)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Decoding error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = v.storage.UpdatePost(postForUpdate); err != nil {
		log.Printf("Update error: %s", err.Error())
		v.ReturnErrorResponse(NewDecodingError(), w)
		return
	}

	v.ReturnResponse(w, []byte("OK"))
}

// DeletePost
//
//	@Summary		Delete post
//	@Description	Delete single post
//	@Produce		json
//	@Success		200	{array}		models.Post
//	@Failure		400	{object}	V1.JsonError	"if we cannot decode or encode payload"
//	@Failure		500	{object}	V1.InputError	"if we have bad payload"
//	@Router			/v1/post/delete/{ID} [get]
func (v *V1) DeletePost(w http.ResponseWriter, r *http.Request) {

	// Get params from context
	ctx := r.Context()
	var params router.Params
	if pr, ok := ctx.Value("params").(router.Params); ok {
		params = pr
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		v.ReturnErrorResponse(NewInputError(), w)
	}

	if err := v.storage.DeletePost(id); err != nil {
		v.ReturnErrorResponse(NewInputError(), w)
	}

	v.ReturnResponse(w, []byte("OK"))
}

// CreatePost
//
//	@Summary		Create post
//	@Description	Create a post
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/V1/post/create-post [post]
func (v *V1) CreatePost(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Problem: %v", err)
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

// CreateComment
//
//	@Summary		Create post
//	@Description	Create a post
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/V1/post/create-comment [post]
func (v *V1) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Decoding error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := v.storage.CreateComment(&comment); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("created"))
	log.Printf("Comment with id [%d] was created", comment.ID)
}

func (v *V1) getSizeAndOffset(params router.Params) (int, int, error) {
	size, err := strconv.Atoi(params["size"])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot get size: %w", err)
	}

	indent, err := strconv.Atoi(params["indent"])
	if err != nil {
		return 0, 0, fmt.Errorf("cannot get offset: %w", err)
	}

	return size, indent, nil
}
