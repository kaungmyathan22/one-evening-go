package main

import "fmt"

func Alphabet(length int) []string {
	str := []string{}
	for i := 0; i < length; i++ {
		str = append(str, characterByIndex(i))
	}
	return str
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
