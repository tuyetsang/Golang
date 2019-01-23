package models

import (
	"database/sql"
	"entities"
)

type ProductModelss7 struct {
	Db *sql.DB
}

func (productModelss7 ProductModelss7) FindAll() ([]entities.Productss7, error) {
	rows, err := productModelss7.Db.Query("select * from product")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Productss7
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int
			rows.Scan(&id, &name, &price, &quantity)
			products = append(products, entities.Productss7{
				Id:       id,
				Name:     name,
				Price:    price,
				Quantity: quantity,
			})
		}
		return products, nil
	}
}

func (productModelss7 ProductModelss7) Find(id int64) (entities.Productss7, error) {
	rows, err := productModelss7.Db.Query("select * from product where id=?", id)
	if err != nil {
		return entities.Productss7{}, err
	} else {
		var product entities.Productss7
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int
			rows.Scan(&id, &name, &price, &quantity)
			product = entities.Productss7{
				Id:       id,
				Name:     name,
				Price:    price,
				Quantity: quantity,
			}
		}
		return product, nil
	}
}

func (productModelss7 ProductModelss7) Search(min, max int) ([]entities.Productss7, error) {
	rows, err := productModelss7.Db.Query("select * from product where price between ? and ?", min, max)
	if err != nil {
		return nil, err
	} else {
		var products []entities.Productss7
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int
			rows.Scan(&id, &name, &price, &quantity)
			products = append(products, entities.Productss7{
				Id:       id,
				Name:     name,
				Price:    price,
				Quantity: quantity,
			})
		}
		return products, nil
	}
}

func (productModelss7 ProductModelss7) FindKeyWord(key string) ([]entities.Productss7, error) {
	rows, err := productModelss7.Db.Query("select * from product where name like ?", "%"+key+"%")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Productss7
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int
			rows.Scan(&id, &name, &price, &quantity)
			products = append(products, entities.Productss7{
				Id:       id,
				Name:     name,
				Price:    price,
				Quantity: quantity,
			})
		}
		return products, nil
	}
}

func (productModelss7 ProductModelss7) Create(product *entities.Productss7) error {
	result, err := productModelss7.Db.Exec("Insert into product(name,price,quantity) values (?,?,?)", product.Name, product.Price, product.Quantity)
	if err != nil {
		return err
	} else {
		product.Id, _ = result.LastInsertId()
		return nil
	}
}

func (productModelss7 ProductModelss7) Delete(id int) (int64, error) {
	result, err := productModelss7.Db.Exec("Delete from product where id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (productModelss7 ProductModelss7) Update(product entities.Productss7) (int64, error) {
	result, err := productModelss7.Db.Exec("update product set name=?,price=?,quantity=? where id=?", product.Name, product.Price, product.Quantity)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
