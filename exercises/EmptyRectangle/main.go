package main

import "fmt"

func drawRectangle(height, width int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	var height, width int

	fmt.Print("Please, enter height of rectangle: ")
	if _, err := fmt.Scan(&height); err != nil || height <= 0 {
		fmt.Println("Error: please enter a positive integer for height.")
		return
	}

	fmt.Print("Please, enter width of rectangle: ")
	if _, err := fmt.Scan(&width); err != nil || width <= 0 {
		fmt.Println("Error: please enter a positive integer for width.")
		return
	}

	drawRectangle(height, width)
}
