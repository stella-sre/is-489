package tax

import "errors"

var ErrInvalidAmount = errors.New("tax: invalid amount")

const (
	lowThreshold  = 1000.0
	midThreshold  = 10000.0
	lowRate       = 0.0
	midRate       = 0.10
	highRate      = 0.15
)

type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Calculate(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, ErrInvalidAmount
	}
	switch {
	case amount < lowThreshold:
		return amount * lowRate, nil
	case amount <= midThreshold:
		return amount * midRate, nil
	default:
		return amount * highRate, nil
	}
}