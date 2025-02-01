package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Coops struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	ImageURL string `json:"imageUrl"`
}

func GetCoops(db *pgxpool.Pool) ([]Coops, error) {
	query := `
		SELECT
			1 AS ID,
			'agraria' AS name,
			'industry' AS category,
			'Cooperativa industrial.' AS desc,
			'' AS ImageURL
	`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unabled to query coops: %w", err)
	}
	defer rows.Close()

	coops, err := pgx.CollectRows(rows, pgx.RowToStructByName[Coops])
	if err != nil {
		return nil, fmt.Errorf("unabled to collect coops rows: %w", err)
	}

	return coops, nil
}
