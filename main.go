package main

import "fmt"

func main() {
	message, _ := enterTheClub(70)
	fmt.Println(message)
	// fmt.Println(entered)
}

func enterTheClub(age int) (string, bool) {
	if age >= 18 && age < 45 {
		return "Please Enter", true
	} else if age >= 45 && age < 65 {
		return "Maybe you'll change your mind", true
	} else if age >= 65 {
		return "You are so old", true
	}

	return "You are under 18 years old", false
}
