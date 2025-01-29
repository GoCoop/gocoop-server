package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func (c *Categories) GetCategories(db *pgxpool.Pool) (*[]Categories, error) {
	categories := []Categories{} // Get data from database

	return &categories, nil
}
