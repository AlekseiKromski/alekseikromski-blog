package jwt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestJWTGuard_Check(t *testing.T) {
	g := New("SecretYouShouldHide")
	token := g.Auth(1)
	log.Println(token)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	assert.True(t, g.Check(req))
}
