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
	{id: 1, name: "pt-BR"},
	{id: 2, name: "en-US"},
}

func returnLangId(lang string) (id int) {
	for _, l := range Locales {
		if strings.Contains(l.name, lang) {
			return l.id
		}
	}

	return Locales[1].id
}
