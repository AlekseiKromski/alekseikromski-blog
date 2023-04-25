package v1

import (
	"encoding/json"
	"net/http"
)

// GetAllTags
//
//	@Summary		Get all tags
//	@Description	Return all tags, that we have
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/V1/tag/all [get]
func (v *V1) GetAllTags(w http.ResponseWriter, r *http.Request) {

	tags := v.storage.GetTags()
	if len(tags) == 0 {
		v.ReturnErrorResponse(NewInputError(), w)
		return
	}

	response, err := json.Marshal(tags)
	if err != nil {
		v.ReturnErrorResponse(NewDecodingError(), w)
		return
	}

	v.ReturnResponse(w, response)
}
