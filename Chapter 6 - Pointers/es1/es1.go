package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func makePerson(firstname string, lastname string, age int) Person {
	person := Person{
		Firstname: firstname,
		Lastname:  lastname,
		Age:       age,
	}

	return person
}

func makePersonPointer(firstname string, lastname string, age int) *Person {
	return &Person{
		Firstname: firstname,
		Lastname:  lastname,
		Age:       age,
	}
}

func main() {
	p := makePerson("kevin", "koltraka", 12)
	pp := makePersonPointer("mario", "rossi", 34)
	/* Surprisingly, it also says that p escapes to the heap. The reason for this is that it is passed into fmt.Println.
	This is because the parameter to fmt.Println are ...any. The current Go compiler moves to the heap any value that is
	passed in to a function via a parameter that is of an interface type. (Interfaces are covered in Chapter 7.) */
	fmt.Println(p)
	fmt.Println(pp)
}
