package main

import "fmt"

func sumDigitsInNumber(n int) int {
	result := 0
	for n != 0 {
		result += n % 10
		n /= 10
	}
	if result < 0 {
		return -result
	}
	return result
}

func main() {
	var number int
	fmt.Print("Please, enter integer: ")
	if _, err := fmt.Scan(&number); err != nil {
		fmt.Println("Error: invalid input. Please enter an integer number only.")
		return
	}

	fmt.Println("Sum of digits:", sumDigitsInNumber(number))
}
