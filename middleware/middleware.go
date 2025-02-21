package middleware

import (
	"context"
	"net/http"
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

func HandleAcceptLang(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		lang := req.Header.Get("Accept-Language")
		userLang := Locales[1]
		if lang != "" {
			l := accept.Parse(lang)
			lang = l[0].Type
			langId := returnLangId(lang)
			userLang = Locale{Id: langId, Name: lang}
		}
		ctx := context.WithValue(req.Context(), LangKey, userLang)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
