package main

import "strings"

func CountCreatedEvents(events []string) int {
	count := 0
	for _, v := range events {
		if strings.HasSuffix(v, "deleted") {
			break
		} else if strings.HasSuffix(v, "created") {
			count += 1
		}
	}
	return count
}

func main() {
	events := []string{
		"product_created",
		"product_updated",
		"product_assigned",
		"order_created",
		"order_updated",
		"client_created",
		"client_updated",
		"client_refreshed",
		"client_deleted",
		"order_updated",
	}

	CountCreatedEvents(events)
}
