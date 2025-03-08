package middleware

import (
	"net/http"
)

func Wrapper(next http.Handler) http.Handler {
	return CORS(HandleAcceptLang(RateLimit(next)))
}
