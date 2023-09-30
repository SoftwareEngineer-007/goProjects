package main

import (
	"errors"
	"fmt"
)

func main() {
	// message, err := enterTheClub(36)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println(prediction("thue"))
	
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

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "sun":
		return "Today is a day off. Do you work tomorrow!", nil
	case "mon":
		return "Have a good start to the week. Today is Monday", nil
	case "tue":
		return "Have a nice Tuesday", nil
	case "wed":
		return "Have a great Wednesday", nil
	case "thu":
		return "Have a wonderful Thursday", nil
	case "fri":
		return "End of the work week! Today is Friday", nil
	case "sat":
		return "You can relax. Today is Saturday", nil
	default:
		return "Enter the correct day of the week", errors.New("invalid day of week")
	}
}
