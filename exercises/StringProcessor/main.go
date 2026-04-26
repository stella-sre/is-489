package main

import (
	"fmt"
	"strings"
)

const INPUT_DATA = "Login; Name; Email\n" +
	"peterson; Chris Peterson; peterson@outlook.com\n" +
	"james; Derek James; james@gmail.com\n" +
	"jackson; Walter Jackson; jackson@gmail.com\n" +
	"gregory; Mike Gregory; gregory@yahoo.com"

func convert1(input string) string {
	result := ""
	lines := strings.Split(input, "\n")
	for _, line := range lines[1:] {
		parts := strings.Split(line, ";")
		if len(parts) >= 3 {
			result += strings.TrimSpace(parts[0]) + " ==> " + strings.TrimSpace(parts[2]) + "\n"
		}
	}
	return result
}

func convert2(input string) string {
	result := ""
	lines := strings.Split(input, "\n")
	for _, line := range lines[1:] {
		parts := strings.Split(line, ";")
		if len(parts) >= 3 {
			result += strings.TrimSpace(parts[1]) + " (email: " + strings.TrimSpace(parts[2]) + ")\n"
		}
	}
	return result
}

func main() {
	fmt.Println("===== Convert 1 =====")
	fmt.Print(convert1(INPUT_DATA))
	fmt.Println("===== Convert 2 =====")
	fmt.Print(convert2(INPUT_DATA))
}
