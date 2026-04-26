package main

import "testing"

func TestGCD_CommonCase(t *testing.T) {
	result := gcdRecursive(12, 8)
	if result != 4 {
		t.Errorf("expected 4, got %d", result)
	}
}

func TestGCD_CoprimeNumbers(t *testing.T) {
	result := gcdRecursive(3, 5)
	if result != 1 {
		t.Errorf("expected 1, got %d", result)
	}
}

func TestGCD_NegativeNumber(t *testing.T) {
	result := gcdRecursive(-12, 8)
	if result != 4 {
		t.Errorf("expected 4, got %d", result)
	}
}

func TestGCD_WithZero(t *testing.T) {
	result := gcdRecursive(0, 5)
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestGCD_LargeNumbers(t *testing.T) {
	result := gcdRecursive(100, 75)
	if result != 25 {
		t.Errorf("expected 25, got %d", result)
	}
}
