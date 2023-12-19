package repository

import (
	"api/app/entity"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) CreateProduct(product entity.Product) error {
	query := "INSERT INTO products (name, price, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"
	_, err := r.db.Exec(query, product.Name, product.Price)
	return err
}

func (r *ProductRepository) GetProducts(filter string, page int, limit int) ([]entity.Product, error) {
	query := "SELECT id, name, price FROM products WHERE deleted_at IS NULL"

	if filter != "" {
		query += fmt.Sprintf(" AND name LIKE '%%%s%%'", filter)
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, (page-1)*limit)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var p entity.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id int) (entity.Product, error) {
	var product entity.Product

	query := "SELECT id, name, price FROM products WHERE id = ? AND deleted_at IS NULL"

	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (r *ProductRepository) UpdateProduct(product entity.Product) error {
	query := "UPDATE products SET name = ?, price = ?, updated_at = NOW() WHERE id = ?"
	_, err := r.db.Exec(query, product.Name, product.Price, product.ID)
	return err
}

func (r *ProductRepository) DeleteProduct(product entity.Product) error {
	query := "UPDATE products SET updated_at = NOW(), deleted_at = NOW() WHERE id = ?"
	_, err := r.db.Exec(query, product.ID)
	return err
}
