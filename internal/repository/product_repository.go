package repository

import (
	"database/sql"
	"fmt"

	"github.com/andreibissolotti/learning_go/internal/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductReository(connection *sql.DB) ProductRepository {
	return ProductRepository{connection: connection}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product " +
		"(product_name, price) " +
		"VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {

	var product model.Product
	query, err := pr.connection.Prepare("SELECT id, product_name, price FROM product " +
		"WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &product, nil
}

func (pr *ProductRepository) DelProductById(id int) error {
	query, err := pr.connection.Prepare("DELETE FROM product " +
		"WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return err
	}

	_, query_err := query.Exec(id)

	if query_err != nil {
		fmt.Println(query_err)
		return query_err
	}

	return nil
}