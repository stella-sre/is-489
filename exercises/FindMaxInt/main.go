package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findMaxIntInArray(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func convertStringArrayToIntArray(parts []string) ([]int, error) {
	result := make([]int, len(parts))
	for i, s := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, fmt.Errorf("invalid value: %q", s)
		}
		result[i] = n
	}
	return result, nil
}

func main() {
	fmt.Print("Please, enter integer numbers separated by space: ")
	var input string
	fmt.Scanln(&input)

	if strings.TrimSpace(input) == "" {
		fmt.Println("Error: no numbers were entered.")
		return
	}

	parts := strings.Fields(input)
	arr, err := convertStringArrayToIntArray(parts)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Please enter only integer numbers separated by spaces.")
		return
	}

	fmt.Println("*** Initial Array ***")
	fmt.Println(arr)
	fmt.Println("*** Max number in array ***")
	fmt.Println(findMaxIntInArray(arr))
}
