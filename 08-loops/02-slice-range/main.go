package main

func Sum(args ...int) int {
	result := 0
	for _, v := range args {
		result += v
	}
	return result
}

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}
