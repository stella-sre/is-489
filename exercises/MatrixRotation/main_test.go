package main

import "testing"

func TestRotate90(t *testing.T) {
	matrix := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}
	rotate90(matrix)
	// expected: [[3.0, 1.0], [4.0, 2.0]]
	if matrix[0][0] != 3.0 || matrix[0][1] != 1.0 ||
		matrix[1][0] != 4.0 || matrix[1][1] != 2.0 {
		t.Errorf("rotate90 incorrect: got %v", matrix)
	}
}

func TestRotate180(t *testing.T) {
	matrix := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}
	rotate180(matrix)
	// expected: [[4.0, 3.0], [2.0, 1.0]]
	if matrix[0][0] != 4.0 || matrix[0][1] != 3.0 ||
		matrix[1][0] != 2.0 || matrix[1][1] != 1.0 {
		t.Errorf("rotate180 incorrect: got %v", matrix)
	}
}

func TestRotate270(t *testing.T) {
	matrix := [][]float64{
		{1.0, 2.0},
		{3.0, 4.0},
	}
	rotate270(matrix)
	// expected: [[2.0, 4.0], [1.0, 3.0]]
	if matrix[0][0] != 2.0 || matrix[0][1] != 4.0 ||
		matrix[1][0] != 1.0 || matrix[1][1] != 3.0 {
		t.Errorf("rotate270 incorrect: got %v", matrix)
	}
}

func TestRotateMatrix_InvalidMode(t *testing.T) {
	matrix := [][]float64{{1.0, 2.0}, {3.0, 4.0}}
	result := rotateMatrix(matrix, 9)
	if result {
		t.Error("expected false for invalid mode, got true")
	}
}

func TestGenerateMatrix_Size(t *testing.T) {
	matrix := generateMatrix(3)
	if len(matrix) != 3 || len(matrix[0]) != 3 {
		t.Errorf("expected 3x3 matrix, got %dx%d", len(matrix), len(matrix[0]))
	}
}
