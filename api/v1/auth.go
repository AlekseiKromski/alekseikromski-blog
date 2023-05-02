package v1

import (
	"encoding/json"
	"log"
	"net/http"
)

type Credits struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login
//
//	@Summary		Login
//	@Description	Login
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/V1/login [post]
func (v *V1) Login(w http.ResponseWriter, r *http.Request) {
	credits := &Credits{}

	err := json.NewDecoder(r.Body).Decode(&credits)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Decoding error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if credits.Username == "" || credits.Password == "" {
		log.Println("Auth error: bad credits")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Find user in db + compare password
	token := v.guards["JWTGuard"].Auth(1)
	v.ReturnResponse(w, []byte(token))
}
