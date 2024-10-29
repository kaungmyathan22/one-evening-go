package main

import "fmt"

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func Keys(products map[int]string) []int {
	keys := make([]int, 0, len(products))
	for product := range products {
		keys = append(keys, product)
	}
	return keys
}

func Values(products map[int]string) []string {
	values := make([]string, 0, len(products))
	for _, v := range products {
		values = append(values, v)
	}
	return values
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}
