package models

import "github.com/jackc/pgx/v5/pgxpool"

type CoopDetails struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ImageUrl   string `json:"imageUrl"`
	Category   string `json:"category"`
	ShortDesc  string `json:"shortDesc"`
	Desc       string `json:"desc"`
	Location   string `json:"location"`
	WebsiteURL string `json:"websiteURL"`
	Workers    int    `json:"workers"`
}

func (c *CoopDetails) GetCoopDetails(db *pgxpool.Pool) (*[]CoopDetails, error) {
	details := []CoopDetails{} // Get data from database

	return &details, nil
}
