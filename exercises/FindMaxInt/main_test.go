package main

import "testing"

func TestFindMaxInt_PositiveNumbers(t *testing.T) {
	arr := []int{3, 17, 5, 2, 99, 41}
	result := findMaxIntInArray(arr)
	if result != 99 {
		t.Errorf("expected 99, got %d", result)
	}
}

func TestFindMaxInt_NegativeNumbers(t *testing.T) {
	arr := []int{-10, -3, -50, -1}
	result := findMaxIntInArray(arr)
	if result != -1 {
		t.Errorf("expected -1, got %d", result)
	}
}

func TestFindMaxInt_MixedNumbers(t *testing.T) {
	arr := []int{-4, 5, 6, 8, 2}
	result := findMaxIntInArray(arr)
	if result != 8 {
		t.Errorf("expected 8, got %d", result)
	}
}

func TestFindMaxInt_SingleElement(t *testing.T) {
	arr := []int{42}
	result := findMaxIntInArray(arr)
	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestConvertStringArrayToIntArray_InvalidInput(t *testing.T) {
	_, err := convertStringArrayToIntArray([]string{"1", "abc", "3"})
	if err == nil {
		t.Error("expected error for invalid input, got nil")
	}
}
