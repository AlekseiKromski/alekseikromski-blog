package v1

import (
	"alekseikromski.com/blog/router"
	"encoding/json"
	"net/http"
	"strconv"
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
	tags := v.storage.GetTags(nil)
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

// DeleteTag
//
//	@Summary		Delete tag
//	@Description	Delete single tag
//	@Produce		json
//	@Success		200	{array}		models.Post
//	@Failure		400	{object}	V1.JsonError	"if we cannot decode or encode payload"
//	@Failure		500	{object}	V1.InputError	"if we have bad payload"
//	@Router			/v1/tag/delete/{ID} [get]
func (v *V1) DeleteTag(w http.ResponseWriter, r *http.Request) {

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

	if err := v.storage.DeleteTag(id); err != nil {
		v.ReturnErrorResponse(NewInputError(), w)
		return
	}

	v.ReturnResponse(w, []byte("OK"))
}
