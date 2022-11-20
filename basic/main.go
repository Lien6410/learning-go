package main

import (
	"log"
	"sort"
	"time"
)

// var birthDate time.Time

type User struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate time.Time
}

func main() {
	// fmt.Println("Hi there.")

	// var whatToSay string

	// whatToSay = "Goodbye, cruel world."

	// fmt.Println(whatToSay)

	// var i int

	// i = 7

	// fmt.Println("i is set to", i)

	// whatWasSaid, andElse := saySomething()

	// fmt.Println("The function returned", whatWasSaid, "and", andElse)

	// myString := "green"

	// log.Println("myString is set to", myString)

	// changePointer(&myString)

	// log.Println("myString is set to", myString)

	// var s = "seven"

	// log.Println("s is set to", s)

	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       16,
	}

	// log.Println(user.FirstName)

	log.Println(user)

	// myMap := make(map[string]string)
	// myMap["firstName"] = "Bob"

	// log.Println(myMap)

	var mySlice []int

	mySlice = append(mySlice, 2, 1, 3)
	sort.Ints(mySlice)

	log.Println(mySlice)
}

func saySomething() (string, string) {
	return "something", "else"
}

func changePointer(s *string) {
	newValue := "red"
	*s = newValue
}
