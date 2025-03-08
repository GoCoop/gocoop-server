package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/timewasted/go-accept-headers"
)

type contextKey string

const LangKey contextKey = "language"

type Locale struct {
	Id   int
	Name string
}

var Locales = []Locale{
	{Id: 1, Name: "pt-br"},
	{Id: 2, Name: "en-us"},
}

func returnLangId(lang string) (id int) {
	lang = strings.ToLower(lang)
	for _, l := range Locales {
		if strings.Contains(l.Name, lang) {
			return l.Id
		}
	}

	return Locales[1].Id
}

func HandleAcceptLang(w http.ResponseWriter, req *http.Request) context.Context {
	lang := req.Header.Get("Accept-Language")
	userLang := Locales[1]

	if lang != "" {
		l := accept.Parse(lang)
		langId := returnLangId(l[0].Type)
		userLang = Locale{Id: langId, Name: lang}
	}

	ctx := context.WithValue(req.Context(), LangKey, userLang)
	return ctx
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		ALLOWED_ORIGIN := os.Getenv("ALLOWED_ORIGIN")

		w.Header().Set("Access-Control-Allow-Origin", ALLOWED_ORIGIN)
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		ctx := HandleAcceptLang(w, req)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
