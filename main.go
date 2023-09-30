package main

import (
	"errors"
	"fmt"
)

func main() {
	message, err := enterTheClub(36)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(message)
	// fmt.Println(entered)
}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Please Enter", nil
	} else if age >= 45 && age < 65 {
		return "Maybe you'll change your mind", nil
	} else if age >= 65 {
		return "You are so old", errors.New("you are so old")
	}

	return "You are under 18 years old", errors.New("you are so young")
}
