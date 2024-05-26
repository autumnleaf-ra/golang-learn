package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Mac Book Pro","price":200000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P001",
		"name":  "Apple Mac Book Pro",
		"price": 2000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
