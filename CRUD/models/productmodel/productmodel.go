package productmodel

import (
	"crud/config"
	"crud/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query("SELECT products.id, categories.name AS category_name, products.name, products.stock, products.description, products.created_at, products.updated_at FROM products JOIN categories ON products.category_id = categories.id")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.Id, &product.Category.Name, &product.Name, &product.Stock, &product.Description, &product.CreatedAt, &product.UpdatedAt); err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec(
		"INSERT INTO products (category_id, name, stock, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		product.Category.Id, product.Name, product.Stock, product.Description, product.CreatedAt, product.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow("SELECT products.id, categories.name AS category_name, products.name, products.stock, products.description, products.created_at, products.updated_at FROM products JOIN categories ON products.category_id = categories.id WHERE products.id = ?", id)

	var product entities.Product
	if err := row.Scan(&product.Id, &product.Category.Name, &product.Name, &product.Stock, &product.Description, &product.CreatedAt, &product.UpdatedAt); err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec("UPDATE products SET category_id = ?, name = ?, stock = ?, description = ?, updated_at = ? WHERE products.id = ?",
	product.Category.Id,
    product.Name,
    product.Stock,
    product.Description,
    product.UpdatedAt,
    id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM products WHERE products.id = ?", id)
	return err
}
