package v1

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type Credits struct {
	Email    string `json:"email"`
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

	if credits.Email == "" || credits.Password == "" {
		log.Println("Auth error: bad credits")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := v.storage.GetUser(credits.Email)
	if err != nil {
		v.ReturnErrorResponse(err, w)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credits.Password)); err != nil {
		log.Printf("password are not the same: %v", err)
		v.ReturnErrorResponse(NewAuthError("bad credits"), w)
		return
	}

	//Find user in db
	token := v.guards["JWTGuard"].Auth(user.ID)
	v.ReturnResponse(w, []byte(token))
}
