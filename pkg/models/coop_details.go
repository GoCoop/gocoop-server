package models

type categoriesData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Icon  string `json:"icon"`
}

type CoopDetails struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	Categories  []categoriesData `json:"categories"`
	ImageURL    string           `json:"image_url"`
	WebsiteURL  string           `json:"website_url"`
	Workers     int              `json:"workers"`
	ShortDesc   string           `json:"short_desc"`
	Description string           `json:"description"`
	Country     string           `json:"country"`
}

type DetailsParams struct {
	Slug   string
	LangId int
}
