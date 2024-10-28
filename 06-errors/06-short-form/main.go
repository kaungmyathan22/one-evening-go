package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("analysis.xlsx"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("ok")
	}
}
