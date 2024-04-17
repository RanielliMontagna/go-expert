package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	// db.Create(&Product{Name: "Laptop", Price: 1000})
	products := []Product{
		{Name: "Mouse", Price: 10},
		{Name: "Keyboard", Price: 20},
		{Name: "Monitor", Price: 200},
	}

	db.Create(&products)

	// Select one
	var product Product
	db.First(&product, 1)
	db.First(&product, "name = ?", "Mouse")
	fmt.Println(product)

	// Select all
	db.Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	// db.Limit(2).Offset(2).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// Where
	db.Where("price > ?", 50).Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	db.Where("name LIKE ?", "%o%").Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	// Update and Delete
	var productToDelete Product
	db.First(&productToDelete, 1)
	productToDelete.Name = "Deleted Product"
	db.Save(&productToDelete)

	db.Delete(&productToDelete)
}
