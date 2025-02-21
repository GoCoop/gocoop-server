package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoriesData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type CoopDetails struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	Categories  []CategoriesData `json:"categories"`
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

func (c *CoopDetails) GetCoopDetails(db *pgxpool.Pool, params DetailsParams) (CoopDetails, error) {

	query := `
		SELECT
			c.id,
			cd.name,
			JSONB_AGG(JSONB_BUILD_OBJECT('id', cat.id, 'name', catT.name, 'icon', cat.name)) AS categories,
			cd.image_url as image_url,
			cd.website_url AS website_url,
			cd.workers AS workers,
			cdt.short_desc,
			cdt.description,
			cdt.country
		FROM t_coops c
		JOIN t_coop_details cd ON
			c.id = cd.id
		JOIN t_coops_categories cc ON
			cc.coop_id = c.id
		JOIN t_categories cat ON
			cat.id = cc.category_id
		LEFT JOIN t_categories_translations catT ON
			catT.category_id = cc.category_id
			AND catT.language_id = $2
		JOIN t_coop_details_translations cdt ON
			cd.id = cdt.coop_id
		WHERE
			c.slug = $1 
			AND cdt.language_id = $2
		GROUP BY
			c.id, 
			cd.name, 
			cd.image_url, 
			cd.website_url, 
			cd.workers,
			cdt.short_desc,
			cdt.description,
			cdt.country;
		`
	rows, err := db.Query(context.Background(), query, params.Slug, params.LangId)
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
