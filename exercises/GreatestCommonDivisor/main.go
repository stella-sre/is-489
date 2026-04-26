package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func gcdRecursive(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}
	return gcdRecursive(b, a%b)
}

func main() {
	fmt.Print("Please, enter two numbers separated by space: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Error: no input was entered.")
		return
	}

	parts := strings.Fields(input)
	if len(parts) != 2 {
		fmt.Println("Error: please enter exactly two numbers separated by a space.")
		return
	}

	a, err1 := strconv.Atoi(parts[0])
	b, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Error: invalid value. Please enter integer numbers only.")
		return
	}

	fmt.Printf("GCD of %d and %d is: %d\n", a, b, gcdRecursive(a, b))
}
