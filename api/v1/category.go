package v1

import (
	"encoding/json"
	"net/http"
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
