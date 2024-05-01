package entity

import (
	"testing"

	"github.com/tj/assert"
)

func TestNewProduct(t *testing.T) {
	t.Run("create new product", func(t *testing.T) {
		name := "product"
		price := 1000.0

		p, err := NewProduct(name, price)

		assert.Nil(t, err)
		assert.NotNil(t, p)
		assert.NotEmpty(t, p.ID)
		assert.Equal(t, name, p.Name)
		assert.Equal(t, price, p.Price)
	})

	t.Run("create new product when name is empty", func(t *testing.T) {
		name := ""
		price := 1000.0

		p, err := NewProduct(name, price)

		assert.Nil(t, p)
		assert.Equal(t, ErrorNameIsRequired, err)
	})

	t.Run("create new product when price is 0", func(t *testing.T) {
		name := "product"
		price := 0.0

		p, err := NewProduct(name, price)

		assert.Nil(t, p)
		assert.Equal(t, ErrorPriceIsRequired, err)
	})

	t.Run("create new product when price is negative", func(t *testing.T) {
		name := "product"
		price := -1000.0

		p, err := NewProduct(name, price)

		assert.Nil(t, p)
		assert.Equal(t, ErrorInvalidPrice, err)
	})

	t.Run("validate product", func(t *testing.T) {
		p := &Product{
			Name:  "product",
			Price: 1000,
		}

		err := p.Validate()

		assert.Nil(t, err)
	})
}
