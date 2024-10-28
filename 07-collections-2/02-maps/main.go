package main

import "fmt"

var (
	Stats = map[string]int{}
)

func CreateUser(user string) {
	count, ok := Stats["create"]
	if !ok {
		Stats["create"] = 1
	} else {
		Stats["create"] = count + 1
	}
	fmt.Println("Creating user", user)
}

func UpdateUser(user string) {
	count, ok := Stats["update"]
	if !ok {
		Stats["update"] = 1
	} else {
		Stats["update"] = count + 1
	}
	fmt.Println("Updating user", user)
}

func PurgeStats() {
	Stats["update"] = 0
	Stats["create"] = 0
}
