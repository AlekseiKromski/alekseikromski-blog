package guard

import "net/http"

type Guard interface {
	Check(req *http.Request) bool
	Auth(userID int) string
}
