package repository

import (
	"database/sql"
	"errors"

	"github.com/prashant-bhilwar/order-processing-system/product-service/database"
	"github.com/prashant-bhilwar/order-processing-system/product-service/model"
)

func CreateProduct(p *model.Product) error {
	query := `INSERT INTO products (name, description, price, in_stock)
			  VALUES ($1, $2, $3, $4)
			  RETURNING id`
	err := database.DB.QueryRow(query, p.Name, p.Description, p.Price, p.InStock).Scan(&p.ID)
	return err
}

func GetAllProducts() ([]model.Product, error) {
	query := `SELECT id, name, description, price, in_stock FROM products`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.InStock)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductbyID(id int) (*model.Product, error) {
	query := `SELECT id, name, price FROM products WHERE id = $1`
	var product model.Product
	err := database.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}
