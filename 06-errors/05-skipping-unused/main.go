package main

import (
	"fmt"
	"os"
)

func CheckFile(str string) bool {
	_, err := os.ReadFile(str)
	return err == nil
}

func main() {
	ok := CheckFile("input.csv")
	if ok {
		fmt.Println("File correctly read")
	} else {
		fmt.Println("Failed to read file")
	}
}
