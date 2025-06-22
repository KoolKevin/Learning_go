package main

import (
	// devo importare esplicitamente il package pippo anche se sono nello stesso modulo
	"es1/pippo"
	"fmt"

	ch10 "github.com/learning-go-book-2e/ch10_solution"
)

// commento
func main() {
	fmt.Println(ch10.Add(10, 20))

	pippo.Pippo()
}
