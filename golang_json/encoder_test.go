package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	w, _ := os.Create("customerout.json")
	encoder := json.NewEncoder(w)

	customer := Customer{
		FirstName: "Muhammad",
		LastName:  "Ramadhan",
	}

	encoder.Encode(customer)
}
