package main

import "fmt"

func main() {
	rect := Rectangle{
		Width:  10,
		Length: 5,
	}

	area := Area(rect)

	fmt.Println("Area:", area)
}

type Rectangle struct {
	Width  int
	Length int
}

func Area(rect Rectangle) int {
	return rect.Length * rect.Width
}
