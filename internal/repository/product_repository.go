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

// Get All product
func (r *ProductRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	query := `
		SELECT id, sku, name, unit, min_stock, created_at, updated_at 
		FROM products
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var p model.Product

		err := rows.Scan(
			&p.ID,
			&p.SKU,
			&p.Name,
			&p.Unit,
			&p.MinStock,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

// Find product by ID
func (r *ProductRepository) GetByID(ctx context.Context, id string) (*model.Product, error) {
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

// Create Product
func (r *ProductRepository) Create(ctx context.Context, p *model.Product) (*model.Product, error) {
	query := `
		INSERT INTO products(sku, name, unit, min_stock)
		VALUES ($1, $2, $3, $4)
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

// Update Product
func (r *ProductRepository) Update(ctx context.Context, p *model.Product) (*model.Product, error) {
	query := `
		UPDATE products
		SET sku = $1, 
			name = $2, 
			unit = $3, 
			min_stock = $4, 
			updated_at = NOW()
		WHERE id = $5
		RETURNING id, sku, name, unit, min_stock, created_at, updated_at
	`
	err := r.db.QueryRowContext(
		ctx,
		query,
		p.SKU,
		p.Name,
		p.Unit,
		p.MinStock,
		p.ID,
	).Scan(
		&p.ID,
		&p.SKU,
		&p.Name,
		&p.Unit,
		&p.MinStock,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

// Delete Product
func (r *ProductRepository) Delete(ctx context.Context, id string) (*model.Product, error) {
	query := `
		DELETE FROM products
		WHERE id = $1
		RETURNING id, sku, name, unit, min_stock, created_at, updated_at
	`

	var p model.Product

	err := r.db.QueryRowContext(ctx, query, id).Scan(
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
