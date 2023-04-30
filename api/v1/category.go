package v1

import (
	"alekseikromski.com/blog/router"
	"encoding/json"
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
	if len(categories) == 0 {
		v.ReturnErrorResponse(NewInputError(), w)
		return
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
