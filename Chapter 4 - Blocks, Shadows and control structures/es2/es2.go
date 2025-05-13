package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intSlice := make([]int, 0, 100)

	for i := 1; i < 100; i++ {
		intSlice = append(intSlice, rand.Intn(100))
	}

outer:
	for _, v := range intSlice {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Sei!")
			break outer
		case v%2 == 0:
			fmt.Println("Due!")
		case v%3 == 0:
			fmt.Println("Tre!")
		default:
			fmt.Println("Nevermind...")
		}

	}

}
