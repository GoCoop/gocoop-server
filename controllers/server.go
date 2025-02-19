package controllers

import (
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	DB *pgxpool.Pool
}

type locale struct {
	id   int
	name string
}

var Locales = []locale{
	{id: 1, name: "pt-br"},
	{id: 2, name: "en-us"},
}

func returnLangId(lang string) (id int) {
	lang = strings.ToLower(lang)
	for _, l := range Locales {
		if strings.Contains(l.name, lang) {
			return l.id
		}
	}

	return Locales[1].id
}
