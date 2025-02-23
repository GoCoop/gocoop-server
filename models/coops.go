package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

func GetCoops(db *pgxpool.Pool, params SearchParams) ([]Coops, error) {
	query := `
		SELECT 
			c.id,
			c.slug,
			cd.name,
			JSONB_AGG(JSONB_BUILD_OBJECT('id', cat.id, 'name', cat.name)) AS categories,
			MAX(cd.image_url) AS image_url,
			MAX(cdt.short_desc) AS short_desc
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
			(c.slug LIKE '%' || $1 || '%' AND cat.name LIKE '%' || $2 || '') 
			AND cdt.language_id = $3
		GROUP by c.id, c.slug, cd.name;`

	rows, err := db.Query(
		context.Background(),
		query,
		params.Query,
		params.Category,
		params.LangId,
	)

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
