package models

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Coops struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	ImageUrl string `json:"imageUrl"`
}

func (c *Coops) GetCoops(db *pgxpool.Pool) (*[]Coops, error) {
	log.Println("> GET request to /coops")

	coops := []Coops{} // Get data from database

	return &coops, nil
}
