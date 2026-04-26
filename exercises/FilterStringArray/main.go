package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func filterWordsByLength(minLength int, words []string) []string {
	var result []string
	for _, w := range words {
		if len(w) >= minLength {
			result = append(result, w)
		}
	}
	if result == nil {
		return []string{}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please, enter any words separated by space: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Error: no words were entered.")
		return
	}

	var minLength int
	fmt.Print("Please, enter minimum word length to filter words: ")
	if _, err := fmt.Scan(&minLength); err != nil || minLength <= 0 {
		fmt.Println("Error: please enter a positive integer for minimum length.")
		return
	}

	words := strings.Fields(input)
	filtered := filterWordsByLength(minLength, words)

	if len(filtered) == 0 {
		fmt.Printf("No words matched the minimum length of %d.\n", minLength)
	} else {
		fmt.Println(filtered)
	}
}
