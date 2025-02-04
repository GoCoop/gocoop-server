package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Coops struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	ShortDesc string `json:"short_desc"`
	ImageURL  string `json:"image_url"`
}

func GetCoops(db *pgxpool.Pool) ([]Coops, error) {
	query := `
		SELECT
			c.id,
			c.slug,
			cd.name,
			cat.name AS category,
			cd.image_url,
			cdt.short_desc
		FROM t_coops c 
		JOIN t_coop_details cd ON
			cd.id = c.id
		JOIN t_coop_details_translations cdt ON
			cdt.coop_id = cd.id
		JOIN t_coops_categories cc ON
			cc.coop_id = cd.id
		JOIN t_categories cat ON
			cat.id = cc.category_id
		WHERE
			c.slug LIKE '%agraria%'
			AND cdt.language_id = 2
			AND cat.id = 1
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
