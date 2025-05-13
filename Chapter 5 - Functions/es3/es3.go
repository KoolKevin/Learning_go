package main

import "fmt"

type prefixAdder func(string) string

func prefixer(prefix string) prefixAdder {
	return func(prefixed string) string {
		return prefix + prefixed
	}
}

func main() {
	p := prefixer("123")
	fmt.Println(p("ciao"))
}
