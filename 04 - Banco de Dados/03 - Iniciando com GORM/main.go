package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"` 
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// Criando um novo produto
	createOneProduct(db)

	// Criando vários novos produtos
	createManyProducts(db)

}

func createOneProduct(db *gorm.DB) error {
	db.Create(&Product{Name: "Notebook", Price: 100.00})
	return nil
}

func createManyProducts(db *gorm.DB) error {
	products := []Product{
		{Name: "Notebook", Price: 100.00},
		{Name: "Notebook", Price: 200.00},
		{Name: "Notebook", Price: 300.00},
	}
	db.Create(&products)
	return nil
}
