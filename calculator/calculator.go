package calculator

import (
	"errors"
	"math"
	)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}


func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of zero or negative numbers not allowed")
	}
	return math.Sqrt(a), nil
}

func AddMany(inputs ...float64) float64 {
	var accum float64
	for _, input := range inputs {
		accum += input
	}
	return accum
}	

func SubtractMany(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	result := inputs[0]
	for _, i := range inputs[1:] {
		result -= i
	}
	return result
}
