package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	point := Point{
		X: 10,
		Y: 5,
	}

	fmt.Println("A point:", point)
}
