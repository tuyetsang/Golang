package models

import (
	"database/sql"
	"entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select*from product")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			rows.Scan(&id, &name, &price, &quantity)
			products = append(products, entities.Product{
				Id:       id,
				Name:     name,
				Price:    price,
				Quantity: quantity,
			})
		}
		return products, nil
	}
}
func (productModel ProductModel) Find(id int64) (entities.Product, error) {
	rows, err := productModel.Db.Query("select*from product where id=?", id)
	if err != nil {
		return entities.Product{}, err
	} else {
		var product entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			rows.Scan(&id, &name, &price, &quantity)
			product.Id = id
			product.Name = name
			product.Price = price
			product.Quantity = quantity
		}
		return product, nil
	}
}
func (productModel ProductModel) Create(product *entities.Product) error {
	result, err := productModel.Db.Exec("insert into product(name,price,quantity) values(?,?,?)", product.Name, product.Price, product.Quantity)
	if err != nil {
		return err
	} else {
		product.Id, _ = result.LastInsertId()
		return nil
	}
}
func (productModel ProductModel) Delete(id int64) (int64, error) {
	result, err := productModel.Db.Exec("delete from product where id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
func (productModel ProductModel) Update(product entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("update product set name=?,price=?,quantity=?", product.Name, product.Price, product.Quantity)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
