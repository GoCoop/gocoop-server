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

func GetCategories(db *pgxpool.Pool, langId int) ([]Categories, error) {

	query := `
		SELECT 
			c.id,
			ct.name,
			c.name AS icon
		FROM t_categories c
		JOIN t_categories_translations ct ON
			c.id = ct.category_id
		WHERE
			ct.language_id = $1
		ORDER BY ct.name;
	`

	rows, err := db.Query(context.Background(), query, langId)
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
