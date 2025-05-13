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

	fmt.Println(intSlice)
}
