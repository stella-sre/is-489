package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func firstCharToTitleCase(text string) string {
	runes := []rune(strings.ToLower(text))
	found := false
	for i, r := range runes {
		if !found && unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			found = true
		} else if unicode.IsSpace(r) {
			found = false
		}
	}
	return string(runes)
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

	fmt.Println(firstCharToTitleCase(input))
}
