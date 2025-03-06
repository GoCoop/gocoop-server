package models

type categories struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Coops struct {
	ID         int          `json:"id"`
	Slug       string       `json:"slug"`
	Name       string       `json:"name"`
	Categories []categories `json:"categories"`
	ShortDesc  string       `json:"short_desc"`
	ImageURL   string       `json:"image_url"`
}

type SearchParams struct {
	Query    string
	Category string
	LangId   int
}
