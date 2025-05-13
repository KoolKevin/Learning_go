package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	// sub1 := greetings[:2]
	// sub2 := greetings[1:4]
	// sub3 := greetings[3:]

	// oppure con la full slice expression
	sub1 := greetings[:2:2]
	sub2 := greetings[1:4:4]
	sub3 := greetings[3:5:5]
	fmt.Println(sub1)
	fmt.Println(sub2)
	fmt.Println(sub3)
}
