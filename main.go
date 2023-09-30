package main

import "fmt"

func main() {
	message := "Hello, World!"
	fmt.Println(message)

	var number int = 16
	var mess string = "End of program"
	var num float32 = 23.582
	var a rune = 'S'
	x := 15
	y := 6

	fmt.Println(number)
	fmt.Println(x + y)
	fmt.Println(number - y)
	fmt.Println(num * float32(y))
	fmt.Println(x / y)
	fmt.Println(a)
	fmt.Println(mess)
}
