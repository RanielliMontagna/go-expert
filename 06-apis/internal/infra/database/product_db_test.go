package database

import (
	"fmt"
	"goexpert/apis/internal/entity"
	"math/rand"
	"testing"

	"github.com/tj/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProductDB(t *testing.T) {
	t.Run("create new product", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		db.AutoMigrate(entity.Product{})

		product, _ := entity.NewProduct("Product 1", 1000)
		productDB := NewProduct(db)

		err = productDB.Create(product)
		assert.Nil(t, err)

		var productFound entity.Product
		err = db.First(&productFound, "id = ?", product.ID).Error
		assert.Nil(t, err)
		assert.Equal(t, product.ID, productFound.ID)
		assert.Equal(t, product.Name, productFound.Name)
		assert.Equal(t, product.Price, productFound.Price)
	})

	t.Run("find all products", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		db.AutoMigrate(entity.Product{})

		for i := 1; i < 24; i++ {
			product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*1000)
			assert.NoError(t, err)
			db.Create(product)
		}

		productDB := NewProduct(db)
		products, err := productDB.FindAll(1, 10, "asc")

		assert.NoError(t, err)
		assert.Len(t, products, 10)
		assert.Equal(t, "Product 1", products[0].Name)
		assert.Equal(t, "Product 10", products[9].Name)

		products, err = productDB.FindAll(2, 10, "asc")
		assert.NoError(t, err)
		assert.Len(t, products, 10)
		assert.Equal(t, "Product 11", products[0].Name)
		assert.Equal(t, "Product 20", products[9].Name)

		products, err = productDB.FindAll(3, 10, "asc")
		assert.NoError(t, err)
		assert.Len(t, products, 3)
		assert.Equal(t, "Product 21", products[0].Name)
		assert.Equal(t, "Product 23", products[2].Name)
	})

	t.Run("find product by id", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		db.AutoMigrate(entity.Product{})

		product, err := entity.NewProduct("Product 1", 1000)
		assert.NoError(t, err)
		db.Create(product)

		productDB := NewProduct(db)
		productFound, err := productDB.FindByID(product.ID.String())

		assert.NoError(t, err)
		assert.Equal(t, "Product 1", productFound.Name)
	})

	t.Run("update product", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		db.AutoMigrate(entity.Product{})

		product, err := entity.NewProduct("Product 1", 1000)
		assert.NoError(t, err)
		db.Create(product)

		product.Name = "Product 2"
		product.Price = 2000

		productDB := NewProduct(db)
		err = productDB.Update(product)

		assert.NoError(t, err)

		product, err = productDB.FindByID(product.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, "Product 2", product.Name)
		assert.Equal(t, 2000.0, product.Price)
	})

	t.Run("delete product", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		db.AutoMigrate(entity.Product{})

		product, err := entity.NewProduct("Product 1", 1000)
		assert.NoError(t, err)
		db.Create(product)

		productDB := NewProduct(db)
		err = productDB.Delete(product.ID.String())

		assert.NoError(t, err)

		product, err = productDB.FindByID(product.ID.String())
		assert.Error(t, err)
		assert.Nil(t, product)

	})
}
