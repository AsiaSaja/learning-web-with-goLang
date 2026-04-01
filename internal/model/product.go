package model

import "time"

type Product struct {
	ID        string
	SKU       string
	Name      string
	Unit      string
	MinStock  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
