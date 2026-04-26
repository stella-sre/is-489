package main

import "testing"

func TestExtendArray_DoublesLength(t *testing.T) {
	arr := []int{1, 2, 3}
	result := extendArray(arr)
	if len(result) != len(arr)*2 {
		t.Errorf("expected length %d, got %d", len(arr)*2, len(result))
	}
}

func TestExtendArray_FirstHalfEqualsOriginal(t *testing.T) {
	arr := []int{10, 20, 30}
	result := extendArray(arr)
	for i, v := range arr {
		if result[i] != v {
			t.Errorf("expected result[%d]=%d, got %d", i, v, result[i])
		}
	}
}

func TestExtendArray_SecondHalfIsMultiplied(t *testing.T) {
	arr := []int{10, 20, 30}
	result := extendArray(arr)
	for i, v := range arr {
		expected := v * MULTIPLIER
		if result[len(arr)+i] != expected {
			t.Errorf("expected result[%d]=%d, got %d", len(arr)+i, expected, result[len(arr)+i])
		}
	}
}

func TestExtendArray_SingleElement(t *testing.T) {
	arr := []int{5}
	result := extendArray(arr)
	if len(result) != 2 || result[0] != 5 || result[1] != 10 {
		t.Errorf("expected [5, 10], got %v", result)
	}
}

func TestGenerateRandomArray_CorrectLength(t *testing.T) {
	result := generateRandomArray(6)
	if len(result) != 6 {
		t.Errorf("expected length 6, got %d", len(result))
	}
}
