package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{})

	// // Create Categories
	// category := Category{Name: "Kitchen"}
	// db.Create(&category)

	// category2 := Category{Name: "Electronics"}
	// db.Create(&category2)

	// // Create Product
	// product := Product{
	// 	Name:       "Microwave",
	// 	Price:      1000,
	// 	Categories: []Category{category, category2},
	// }
	// db.Create(&product)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Printf("- Category: %s\n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("-- Product: %s\n", product.Name)
		}
	}
}
