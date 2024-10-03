package main

import "fmt"

// Same as (users []string)
func addUsers(users ...string) {
	for _, user := range users {
		fmt.Println(user)
	}
}

func removeFromSliceReturnUnordered(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func removeFromSliceReturnOrdered(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
