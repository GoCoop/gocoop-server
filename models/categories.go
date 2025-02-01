package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func GetCategories(db *pgxpool.Pool) ([]Categories, error) {

	query := `
		SELECT
			1 					AS id,
			'Indústria' AS name,
			'industry' 	AS icon
		UNION
		SELECT
			2 					AS id,
			'Cerveja' 	AS name,
			'beer' 			AS icon
	`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unabled to query categories: %w", err)
	}
	defer rows.Close()

	categories, err := pgx.CollectRows(rows, pgx.RowToStructByName[Categories])
	if err != nil {
		return nil, fmt.Errorf("unabled to collect rows categories: %w", err)
	}

	return categories, nil
}
