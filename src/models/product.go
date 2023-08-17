package models

import "go-crud/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func GetProducts() []Product {
	db := db.DbConnection()

	selectProducts, err := db.Query("SELECT * FROM product")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func InsertProduct(name, description string, price float64, amount int) {
	db := db.DbConnection()

	insertProduct, err := db.Prepare("INSERT INTO product(name, description, price, amount) VALUES($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, amount)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DbConnection()

	deleteProduct, err := db.Prepare("DELETE FROM product WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.DbConnection()

	dbProduct, err := db.Query("SELECT * FROM product WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	for dbProduct.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = dbProduct.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount
	}
	defer db.Close()
	return p
}

func UpdateProduct(id, name, description string, price float64, amount int) {
	db := db.DbConnection()

	updateProduct, err := db.Prepare("UPDATE product SET name=$1, description=$2, price=$3, amount=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, amount, id)
	defer db.Close()
}
