package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Muhammad",
		LastName:  "Ramadhan",
		Hobbies:   []string{"Game", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","LastName":"Ramadhan","Age":0,"Married":false,"Hobbies":["Game","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	cust := Customer{
		FirstName: "Rama",
		Addresses: []Address{
			{
				Street:     "GG Haji Karim",
				Country:    "Indonesia",
				PostalCode: "1999",
			},
			{
				Street:     "Jl aja",
				Country:    "Indonesia",
				PostalCode: "194949",
			},
		},
	}

	bytes, _ := json.Marshal(cust)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Rama","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"GG Haji Karim","Country":"Indonesia","PostalCode":"1999"},{"Street":"Jl aja","Country":"Indonesia","PostalCode":"194949"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"GG Haji Karim","Country":"Indonesia","PostalCode":"1999"},{"Street":"Jl aja","Country":"Indonesia","PostalCode":"194949"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "GG Haji Karim",
			Country:    "Indonesia",
			PostalCode: "1999",
		},
		{
			Street:     "Jl aja",
			Country:    "Indonesia",
			PostalCode: "194949",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
