package repository

import (
	"context"
	"database/sql"

	"github.com/luqm4n-Al/go-web/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindByID(ctx context.Context, id string) (*model.Product, error) {
	query := `
		SELECT id, sku, name, unit, min_stock, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var p model.Product

	err := row.Scan(
		&p.ID,
		&p.SKU,
		&p.Name,
		&p.Unit,
		&p.MinStock,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func (r *ProductRepository) Create(ctx context.Context, p *model.Product) (*model.Product, error) {
	query := `
		INSERT INTO products(id, sku, name, unit, min_stock)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		p.SKU,
		p.Name,
		p.Unit,
		p.MinStock,
	).Scan(
		&p.ID,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return p, nil

}
