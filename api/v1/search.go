package v1

import (
	"alekseikromski.com/blog/api/storage"
	"encoding/json"
	"log"
	"net/http"
)

// Search
//
//	@Summary		Search
//	@Description	Get all result by search
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/v1/search/query [get]
func (v *V1) Search(w http.ResponseWriter, r *http.Request) {
	var searchRequest *storage.SearchRequest

	err := json.NewDecoder(r.Body).Decode(&searchRequest)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Decoding error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := v.storage.Search(searchRequest)

	payload, err := json.Marshal(result)
	if err != nil {
		v.ReturnErrorResponse(NewDecodingError(), w)
		return
	}

	v.ReturnResponse(w, payload)
}
