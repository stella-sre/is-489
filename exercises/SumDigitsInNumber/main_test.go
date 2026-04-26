package main

import "testing"

func TestSumDigits_PositiveNumber(t *testing.T) {
	result := sumDigitsInNumber(1234)
	if result != 10 {
		t.Errorf("expected 10, got %d", result)
	}
}

func TestSumDigits_NegativeNumber(t *testing.T) {
	result := sumDigitsInNumber(-567)
	if result != 18 {
		t.Errorf("expected 18, got %d", result)
	}
}

func TestSumDigits_Zero(t *testing.T) {
	result := sumDigitsInNumber(0)
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestSumDigits_SingleDigit(t *testing.T) {
	result := sumDigitsInNumber(9)
	if result != 9 {
		t.Errorf("expected 9, got %d", result)
	}
}

func TestSumDigits_LargeNumber(t *testing.T) {
	result := sumDigitsInNumber(99999)
	if result != 45 {
		t.Errorf("expected 45, got %d", result)
	}
}
