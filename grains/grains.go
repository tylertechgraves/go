package grains

import (
	"errors"
	"math"
)

func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("The square number must be greater than 0 and less than 64")
	}

	return uint64(math.Pow(2, float64(square-1))), nil
}

func Total() uint64 {
	value, _ := Square(65)
	return value - 1
}
