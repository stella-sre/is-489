package main

import "fmt"

func generateMatrix(size int) [][]float64 {
	matrix := make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size)
		for j := range matrix[i] {
			matrix[i][j] = float64(i) + float64(j)/10.0
		}
	}
	return matrix
}

func printMatrix(matrix [][]float64) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%.1f\t", val)
		}
		fmt.Println()
	}
}

func transposeMatrix(matrix [][]float64) {
	for i := range matrix {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func verticalReflection(matrix [][]float64) {
	n := len(matrix)
	for i := range matrix {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func horizontalReflection(matrix [][]float64) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
}

func rotate90(matrix [][]float64) {
	transposeMatrix(matrix)
	verticalReflection(matrix)
}

func rotate180(matrix [][]float64) {
	verticalReflection(matrix)
	horizontalReflection(matrix)
}

func rotate270(matrix [][]float64) {
	transposeMatrix(matrix)
	horizontalReflection(matrix)
}

func rotateMatrix(matrix [][]float64, mode int) bool {
	switch mode {
	case 1:
		fmt.Println("90 degrees rotated:")
		rotate90(matrix)
	case 2:
		fmt.Println("180 degrees rotated:")
		rotate180(matrix)
	case 3:
		fmt.Println("270 degrees rotated:")
		rotate270(matrix)
	default:
		fmt.Println("Error: you selected a non-existing mode. Please choose 1, 2 or 3.")
		return false
	}
	return true
}

func main() {
	var size int
	fmt.Print("Please, enter matrix size: ")
	if _, err := fmt.Scan(&size); err != nil || size <= 0 {
		fmt.Println("Error: please enter a positive integer for matrix size.")
		return
	}

	matrix := generateMatrix(size)

	fmt.Println("How do you want to rotate the matrix?")
	fmt.Println("\t1 - 90 degrees")
	fmt.Println("\t2 - 180 degrees")
	fmt.Println("\t3 - 270 degrees")

	var mode int
	if _, err := fmt.Scan(&mode); err != nil {
		fmt.Println("Error: invalid input. Please enter 1, 2 or 3.")
		return
	}

	fmt.Println("\nBase matrix:")
	printMatrix(matrix)
	fmt.Println()

	if rotateMatrix(matrix, mode) {
		printMatrix(matrix)
	}
}
