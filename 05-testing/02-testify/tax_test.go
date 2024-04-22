package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	t.Run("amount greater than or equal to 20000", func(t *testing.T) {
		amount := 20000.0
		expected := 20.0

		result, err := CalculateTax(amount)

		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("amount greater than or equal to 1000", func(t *testing.T) {
		amount := 1000.0
		expected := 10.0

		result, err := CalculateTax(amount)

		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("amount less than 1000", func(t *testing.T) {
		amount := 500.0
		expected := 5.0

		result, err := CalculateTax(amount)

		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("amount less than or equal to 0", func(t *testing.T) {
		amount := 0.0
		result, err := CalculateTax(amount)

		assert.Error(t, err, "amount must be greater than 0")
		assert.Equal(t, 0.0, result)
	})
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}

	repository.On("SaveTax", 10.0).Return(nil).Twice()
	repository.On("SaveTax", 0.0).Return(errors.New("amount must be greater than 0"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "amount must be greater than 0")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 3)
}
