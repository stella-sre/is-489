package main

import (
	"fmt"
	"math/rand"
)

const MULTIPLIER = 2

func generateRandomArray(length int) []int {
	arr := make([]int, length)
	for i := range arr {
		arr[i] = rand.Intn(100) + 1
	}
	return arr
}

func extendArray(arr []int) []int {
	newLength := len(arr) * MULTIPLIER
	result := make([]int, newLength)
	copy(result, arr)
	for i := len(arr); i < newLength; i++ {
		result[i] = arr[i-len(arr)] * MULTIPLIER
	}
	return result
}

func main() {
	var length int
	fmt.Print("Please, enter length of initial array: ")
	if _, err := fmt.Scan(&length); err != nil || length <= 0 {
		fmt.Println("Error: please enter a positive integer for array length.")
		return
	}

	arr := generateRandomArray(length)
	extended := extendArray(arr)

	fmt.Println("*** Initial array ***")
	fmt.Println(arr)
	fmt.Println("*** Extended array ***")
	fmt.Println(extended)
}
