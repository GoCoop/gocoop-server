package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CoopDetails struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ImageURL   string `json:"imageURL"`
	Category   string `json:"category"`
	ShortDesc  string `json:"shortDesc"`
	Desc       string `json:"desc"`
	Location   string `json:"location"`
	WebsiteURL string `json:"websiteURL"`
	Workers    int    `json:"workers"`
}

func (c *CoopDetails) GetCoopDetails(db *pgxpool.Pool, pathV string) (CoopDetails, error) {
	fmt.Println(pathV) // Will be used inside SELECT query

	// Example query
	query := `
		SELECT 
			1 AS id, 
			'Agraria' AS name, 
			'/agraria-logo.jpg' AS imageURL, 
			'industry' AS category, 
			'Cooperativa' AS shortDesc, 
			'desc' AS desc, 
			'Brasil' AS location, 
			'' AS websiteURL,
			400 AS workers`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return CoopDetails{}, fmt.Errorf("unable to query coop details: %w", err)
	}
	defer rows.Close()

	details, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[CoopDetails])
	if err != nil {
		return CoopDetails{}, fmt.Errorf("unable to collect one row coop details: %w", err)
	}

	return details, nil
}
