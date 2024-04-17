package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int `gorm:"foreignKey:ID"`
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{})

	// // Create Category
	// category := Category{Name: "Electronics"}
	// db.Create(&category)

	// // Create Product
	// product := Product{Name: "Laptop", Price: 1000, CategoryID: category.ID}
	// db.Create(&product)

	// product2 := Product{Name: "Mobile", Price: 500, CategoryID: 1}
	// db.Create(&product2)

	var products []Product
	db.Preload("Category").Find(&products)

	for _, product := range products {
		fmt.Printf("Product: %s, Price: %f, Category: %s\n", product.Name, product.Price, product.Category.Name)
	}
}
