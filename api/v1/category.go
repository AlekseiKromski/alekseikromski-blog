package v1

import (
	"alekseikromski.com/blog/api/storage/models"
	"alekseikromski.com/blog/router"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GetAllCategories
//
//	@Summary		Get all categories
//	@Description	Return all categories, that we have
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/V1/category/all [get]
func (v *V1) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	categories := v.storage.GetCategories()
	if categories == nil {
		categories = []*models.Category{}
	}

	response, err := json.Marshal(categories)
	if err != nil {
		v.ReturnErrorResponse(NewDecodingError(), w)
		return
	}

	v.ReturnResponse(w, response)
}

// DeleteCategory
//
//	@Summary		Delete category
//	@Description	Delete single category
//	@Produce		json
//	@Success		200	{array}		models.Post
//	@Failure		400	{object}	V1.JsonError	"if we cannot decode or encode payload"
//	@Failure		500	{object}	V1.InputError	"if we have bad payload"
//	@Router			/v1/category/delete/{ID} [get]
func (v *V1) DeleteCategory(w http.ResponseWriter, r *http.Request) {

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

	if err := v.storage.DeleteCategory(id); err != nil {
		v.ReturnErrorResponse(NewInputError(), w)
	}

	v.ReturnResponse(w, []byte("OK"))
}

// UpdateCategory
//
//	@Summary		Update category
//	@Description	Update single category
//	@Produce		json
//	@Success		200	{array}		models.Post
//	@Failure		400	{object}	V1.JsonError	"if we cannot decode or encode payload"
//	@Failure		500	{object}	V1.InputError	"if we have bad payload"
//	@Router			/v1/category/edit [post]
func (v *V1) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	categoryForUpdate := models.CreateCategory()

	err := json.NewDecoder(r.Body).Decode(&categoryForUpdate)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Decoding error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = v.storage.UpdateCategory(categoryForUpdate); err != nil {
		log.Printf("Update error: %s", err.Error())
		v.ReturnErrorResponse(NewDecodingError(), w)
		return
	}

	v.ReturnResponse(w, []byte("OK"))
}

// CreateCategory
//
//	@Summary		Create category
//	@Description	Create a category
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/V1/category/create [post]
func (v *V1) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Decoding error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	created, err := v.storage.CreateCategory(&category)
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
	log.Printf("Category with id [%d] was created", category.ID)
}
