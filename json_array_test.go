package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

//JSON Primitive Array

func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Syauqi",
		MiddleName: "Damario",
		LastName:   "Djohan",
		Hobbies:    []string{"Cleaning", "Podcast", "Sitcoms"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

//JSON Complex Array
func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Syauqi",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "1106",
			},
			{
				Street:     "Hollywood Boulevard",
				Country:    "United States",
				PostalCode: "778899",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

//Decode JSON Array

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Syauqi","MiddleName":"Damario","LastName":"Djohan","Age":24,"Married":false,"Hobbies":["Cleaning", "Podcast", "Sitcoms"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

