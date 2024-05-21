package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

//struct adalah representasi objek JSON
//atribut object adalah atribut struct

type Customer struct {
	FirstName string
	MiddleName string
	LastName string
	Age int
	Married bool
	Hobbies []string
	Addresses []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Syauqi",
		MiddleName: "Damario",
		LastName: "Djohan",
		Age: 24,
		Married: false,
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}