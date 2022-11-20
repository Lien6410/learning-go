package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	Age       int    `json:"age"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "John",
			 "age":11
		},
		{
			"first_name": "Bob",
			 "age":33
		}
	]`

	var unmarshal []Person
	err := json.Unmarshal([]byte(myJson), &unmarshal)
	if err != nil {
		log.Println("Error unmarshalling json: ", err)
	}

	log.Printf("Unmarshaled: %v", unmarshal)

	// write to json from struct
	var mySlice []Person
	m1 := Person{FirstName: "Alice", Age: 16}
	m2 := Person{FirstName: "Mary", Age: 20}
	mySlice = append(mySlice, m1, m2)

	newJson, err2 := json.MarshalIndent(mySlice, "", "\t")
	if err2 != nil {
		log.Println("Error marshal json: ", err2)
	}

	fmt.Println(string(newJson))
}
