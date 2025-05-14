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

// func makePersonPointer(firstname string, lastname string, age int) *Person {
// 	return &Person{
// 		Firstname: firstname,
// 		Lastname:  lastname,
// 		Age:       age,
// 	}
// }

func main() {
	// var kevins []Person
	kevins := make([]Person, 0)

	for i := 0; i < 10_000_000; i++ {
		kevins = append(kevins, makePerson("kevin", "koltraka", 23))
	}

	fmt.Println(kevins[0])
}
