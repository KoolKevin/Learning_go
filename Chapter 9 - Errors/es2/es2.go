package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		var emptyErr emptyFieldError
		if err != nil {
			if errors.Is(err, ErrInvalidId) {
				fmt.Printf("record %d: %+v error: l'id non è valido\n", count, emp)
			} else if errors.As(err, &emptyErr) {
				fmt.Printf("record %d: %+v manca un campo: %v; msg: %v\n", count, emp, emptyErr.EmptyFieldName, emptyErr.Message)
			} else {
				fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
			}
			continue
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID      = regexp.MustCompile(`\w{4}-\d{3}`)
	ErrInvalidId = errors.New("invalid ID")
)

func ValidateEmployee(e Employee) error {
	if len(e.ID) == 0 {
		return emptyFieldError{
			EmptyFieldName: "ID",
			Message:        "missing ID",
		}
	}
	if !validID.MatchString(e.ID) {
		return ErrInvalidId
	}
	if len(e.FirstName) == 0 {
		return emptyFieldError{
			EmptyFieldName: "FirstName",
			Message:        "missing Firstname",
		}
	}
	if len(e.LastName) == 0 {
		return emptyFieldError{
			EmptyFieldName: "LastName",
			Message:        "missing LastName",
		}
	}
	if len(e.Title) == 0 {
		return emptyFieldError{
			EmptyFieldName: "Title",
			Message:        "missing Title",
		}
	}
	return nil
}

type emptyFieldError struct {
	EmptyFieldName string
	Message        string
}

func (e emptyFieldError) Error() string {
	return e.Message
}
