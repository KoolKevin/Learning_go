package main

import (
	"fmt"
)

func updateSlice(sSlice []string, newValue string) {
	lastIndex := len(sSlice) - 1
	sSlice[lastIndex] = newValue
	fmt.Println(sSlice)
}

func growSlice(sSlice []string, newValue string) {
	sSlice = append(sSlice, newValue)
	fmt.Println(sSlice)
}

func main() {
	sSlice := []string{"io", "sono", "un", "grande"}
	fmt.Println(sSlice)
	updateSlice(sSlice, "scemo")
	fmt.Println(sSlice)

	fmt.Println(sSlice)
	growSlice(sSlice, "scemo")
	fmt.Println(sSlice)
}
