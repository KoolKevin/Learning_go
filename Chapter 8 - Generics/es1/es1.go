package main

import "fmt"

type number interface {
	int | float32 | float64
}

func double[T number](num T) T {
	return 2 * num
}

func main() {
	fmt.Println(double(10))
	fmt.Println(double(10.1))
}
