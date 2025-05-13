package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	e1 := Employee{
		"kevin",
		"koltrka",
		1,
	}

	e2 := Employee{
		firstName: "mario",
		lastName:  "rossi",
		id:        2,
	}

	var e3 Employee
	e3.firstName = "Frank"
	e3.lastName = "roffia"
	e3.id = 3

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}
