package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int `gorm:"foreignKey:ID"`
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// // Create Category
	// category := Category{Name: "Electronics"}
	// db.Create(&category)

	// // Create Product
	// product := Product{Name: "Laptop", Price: 1000, CategoryID: category.ID}
	// db.Create(&product)

	// // Create Serial Number
	// serialNumber := SerialNumber{Number: "123456", ProductID: product.ID}
	// db.Create(&serialNumber)

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)

	// for _, product := range products {
	// 	fmt.Printf("Product: %s, Category: %s, Serial Number: %s\n", product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Printf("- Category: %s\n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("-- Product: %s, Serial Number: %s\n", product.Name, product.SerialNumber.Number)
		}
	}
}
