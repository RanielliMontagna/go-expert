package tax

import (
	"errors"
	"time"
)

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("amount must be greater than 0")
	}

	if amount >= 20000 {
		return 20, nil
	}

	if amount >= 1000 {
		return 10, nil
	}

	return 5, nil
}

func CalculateTax2(amount float64) float64 {
	time.Sleep(1 * time.Millisecond)

	if amount == 0 {
		return 0
	}

	if amount >= 1000 {
		return 10
	}

	return 5
}
