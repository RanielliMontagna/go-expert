package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	t.Run("with amount less than 1000", func(t *testing.T) {
		amount := 500.0
		expected := 5.0

		result := CalculateTax(amount)

		if result != expected {
			t.Errorf("CalculateTax(%f) = %f; want %f", amount, result, expected)
		}
	})

	t.Run("with amount greater than or equal to 1000", func(t *testing.T) {
		amount := 1000.0
		expected := 10.0

		result := CalculateTax(amount)

		if result != expected {
			t.Errorf("CalculateTax(%f) = %f; want %f", amount, result, expected)
		}
	})
}

func TestCalculateTaxBatch(t *testing.T) {
	tests := []struct {
		amount   float64
		expected float64
	}{
		{500.0, 5.0},
		{1000.0, 10.0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := CalculateTax(tt.amount)

			if result != tt.expected {
				t.Errorf("CalculateTax(%f) = %f; want %f", tt.amount, result, tt.expected)
			}
		})
	}
}
