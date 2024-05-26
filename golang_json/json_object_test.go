package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Married   bool
	Hobbies   []string
	Addresses []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Muhammad",
		LastName:  "Ramadhan",
		Age:       23,
		Married:   false,
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}
