package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getWordsAmount(text string) int {
	re := regexp.MustCompile(`[\p{P}\s]+`)
	words := re.Split(strings.TrimSpace(text), -1)
	count := 0
	for _, w := range words {
		if w != "" {
			count++
		}
	}
	return count
}

func main() {
	fmt.Print("Please, enter any text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if strings.TrimSpace(input) == "" {
		fmt.Println("Error: no text was entered.")
		return
	}

	fmt.Println("Amount of words in your text:", getWordsAmount(input))
}
