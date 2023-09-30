package main

import "fmt"

func main() {
	message, _ := enterTheClub(16)
	fmt.Println(message)
	// fmt.Println(entered)
}

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "Please Enter", true
	} else {
		return "Entrance closed", false
	}
}
