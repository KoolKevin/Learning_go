package main

import "fmt"

func main() {
	var total int

	for i := 0; i < 10; i++ {
		// total := total + i
		total = total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}
