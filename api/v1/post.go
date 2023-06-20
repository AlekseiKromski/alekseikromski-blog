package v1

import (
	"alekseikromski.com/blog/api/storage"
	"alekseikromski.com/blog/api/storage/models"
	"alekseikromski.com/blog/router"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var urlRegExp = regexp.MustCompile(`/temp/upload-\d+.(jpg|png|jpeg)`)
var extensionsRegExp = regexp.MustCompile("(png|jpg|jpeg)")

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

	if posts == nil {
		posts = []*models.Post{}
	}

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

	if find, converted, err := findTempImages(postForUpdate.Img); err == nil && find {
		postForUpdate.Img = converted
	} else {
		if err != nil {
			v.ReturnErrorResponse(err, w)
			return
		}
	}

	if find, converted, err := findTempImages(postForUpdate.Description); err == nil && find {
		postForUpdate.Description = converted
	} else {
		if err != nil {
			v.ReturnErrorResponse(err, w)
			return
		}
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
		return
	}

	if err := v.storage.DeletePost(id); err != nil {
		v.ReturnErrorResponse(NewInputError(), w)
		return
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

	if find, converted, err := findTempImages(post.Img); err == nil && find {
		post.Img = converted
	} else {
		if err != nil {
			v.ReturnErrorResponse(err, w)
			return
		}
	}

	if find, converted, err := findTempImages(post.Description); err == nil && find {
		post.Description = converted
	} else {
		if err != nil {
			v.ReturnErrorResponse(err, w)
			return
		}
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
	log.Printf("Post with id [%d] was created", post.ID)
}

func findTempImages(data string) (bool, string, error) {
	//define temp images path in description
	pos := urlRegExp.FindAllString(data, -1)

	if len(pos) != 0 {
		final := data
		for _, path := range pos {

			//Content for old file
			bytes, err := os.ReadFile(filepath.Join(".", path))
			if err != nil {
				log.Printf("[HANDLER] Cannot read file: %v", err)
				return false, "", err
			}

			dest := strings.ReplaceAll(string(path), "/temp/", "./store/images/")

			if err = os.WriteFile(dest, bytes, 0644); err != nil {
				log.Printf("[HANDLER] Cannot write file: %v", err)
				return false, "", err
			}

			final = strings.ReplaceAll(final, string(path), dest[1:len(dest)])
		}
		return true, final, nil
	}

	return false, "", nil
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

	createdComment, err := v.storage.CreateComment(&comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(createdComment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Comment with id [%d] was created", comment.ID)
	v.ReturnResponse(w, response)
}

// UploadFile
//
//	@Summary		Upload [png,jpg,jpeg] image file
//	@Description	Upload files to temp directory (only for png,jpg,jpeg)
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/v1/post/upload [post]
func (v *V1) UploadFile(w http.ResponseWriter, r *http.Request) {
	log.Println("[HANDLER] Starting uploading file to temp dir")

	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("[HANDLER] error during getting file")
		v.ReturnErrorResponse(err, w)
		return
	}

	defer file.Close()
	log.Printf("[HANDLER] Uploaded File: %s", handler.Filename)
	log.Printf("[HANDLER] File Size: %d", handler.Size)
	log.Printf("[HANDLER] MIME Header: %s", handler.Header)

	// Create a temp file within our temp-images directory that follows
	// a particular naming pattern
	pos := extensionsRegExp.FindStringSubmatch(handler.Filename)
	if len(pos) == 0 {
		log.Printf("[HANDLER] incorrect format of file from regexp")
		v.ReturnErrorResponse(fmt.Errorf("incorrect fromat"), w)
		return
	}
	tempFile, err := os.CreateTemp("./temp", fmt.Sprintf("upload-*.%s", pos[0]))
	if err != nil {
		log.Printf("[HANDLER] cannot create temp file: %v", err)
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("[HANDLER] cannot read writed temp file")
		fmt.Println(err)
	}
	// write this byte array to our temp file
	tempFile.Write(fileBytes)

	response, err := json.Marshal(struct {
		Payload string `json:"payload"`
	}{
		Payload: tempFile.Name()[2:len(tempFile.Name())],
	})

	if err != nil {
		log.Printf("[HANDLER] cannot marshal payload %v", err)
		v.ReturnErrorResponse(err, w)
		return
	}
	v.ReturnResponse(w, response)
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
