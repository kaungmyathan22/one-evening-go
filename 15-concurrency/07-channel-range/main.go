package main

import "fmt"

func main() {
	messagesChan := make(chan string)

	go SendMessages(100, messagesChan)

	messages := ReadMessages(messagesChan)
	fmt.Println("Received", len(messages), "messages")
}

func SendMessages(count int, messages chan<- string) {
	for i := 0; i < count; i++ {
		messages <- fmt.Sprintf("message number %v", i+1)
	}
	close(messages)
}

func ReadMessages(messages <-chan string) []string {
	var messagesSlice []string
	for {
		msg, ok := <-messages
		if !ok {
			break
		}

		messagesSlice = append(messagesSlice, msg)
	}

	return messagesSlice
}
