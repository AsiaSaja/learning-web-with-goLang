package service

import (
	"context"
	"errors"

	"github.com/luqm4n-Al/go-web/internal/model"
	"github.com/luqm4n-Al/go-web/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// Service untuk GetAll Product
func (s *ProductService) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	product, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// Service untuk GetById Product
func (s *ProductService) GetByIDProduct(ctx context.Context, id string) (*model.Product, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// Service untuk Create Product
func (s *ProductService) CreateProduct(ctx context.Context, p *model.Product) (*model.Product, error) {
	if p.SKU == "" {
		return nil, errors.New("sku is required")
	}

	if p.Name == "" {
		return nil, errors.New("name is required")
	}

	if p.MinStock < 0 {
		return nil, errors.New("min_stock cannot be negative")
	}

	if p.Unit == "" {
		p.Unit = "pcs"
	}

	product, err := s.repo.Create(ctx, p)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// Service untuk Update Product
func (s *ProductService) UpdateProduct(ctx context.Context, p *model.Product) (*model.Product, error) {
	exist, err := s.repo.GetByID(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	if exist == nil {
		return nil, errors.New("Product Not Found")
	}

	if p.SKU != "" {
		exist.SKU = p.SKU
	}

	if p.Name != "" {
		exist.Name = p.Name
	}

	if p.MinStock >= 0 {
		exist.MinStock = p.MinStock
	}

	if p.Unit != "" {
		exist.Unit = p.Unit
	}

	if exist.SKU == "" {
		return nil, errors.New("sku is required")
	}

	if exist.Name == "" {
		return nil, errors.New("name is required")
	}

	if exist.MinStock < 0 {
		return nil, errors.New("min_stock cannot be negative")
	}

	updated, err := s.repo.Update(ctx, exist)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) (*model.Product, error) {
	return s.repo.Delete(ctx, id)
}
