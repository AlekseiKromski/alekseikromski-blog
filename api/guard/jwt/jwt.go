package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

type JWTGuard struct {
	secret []byte
}

func New(secret string) *JWTGuard {
	return &JWTGuard{
		secret: []byte(secret),
	}
}

// Check - will check JWT token by req
func (jg *JWTGuard) Check(req *http.Request) bool {
	tokenRequest := req.Header.Get("Authorization")
	if tokenRequest == "" || len(tokenRequest) < 10 {
		log.Printf("[JWTGUARD] there is not token in request: %s", req.URL.String())
		return false
	}

	tokenRequest = tokenRequest[7:len(tokenRequest)]

	token, err := jwt.Parse(tokenRequest, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong sign method")
		}
		claims := token.Claims.(jwt.MapClaims)
		if claims["userID"] == nil {
			return nil, fmt.Errorf("wrong format of JWT")
		}
		return jg.secret, nil
	})

	if err != nil {
		log.Printf("[JWTGUARD] token verify failed: %v", err)
		return false
	}

	if token.Valid {
		return true
	}

	return false
}

func (jg *JWTGuard) Auth(userID int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * (time.Hour * 24)).Unix()
	claims["authorized"] = true
	claims["userID"] = userID

	tokenString, err := token.SignedString(jg.secret)
	if err != nil {
		log.Fatalf("cannot create token: %v", err)
	}
	return tokenString
}
