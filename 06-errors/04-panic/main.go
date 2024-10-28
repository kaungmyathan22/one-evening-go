package main

import (
	"errors"
)

var message = ""

func StoreMessage(m string) error {
	if m == "" {
		return errors.New("no message")
	}

	message = m

	return nil
}

func MustStoreMessage(str string) {
	err := StoreMessage(str)
	if err != nil {
		panic(err.Error())
	}
}
func main() {
	MustStoreMessage("Hello!")
}
