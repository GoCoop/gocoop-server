package services

import (
	"context"
	"fmt"
	"gocoop-server/pkg/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCategories(db *pgxpool.Pool, langId int) ([]models.Categories, error) {

	query := `
		SELECT 
			c.id,
			ct.name,
			c.name AS label,
			c.icon
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

	categories, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Categories])
	if err != nil {
		return nil, fmt.Errorf("unabled to collect rows categories: %w", err)
	}

	return categories, nil
}
