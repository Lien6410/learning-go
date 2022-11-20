package main

import "fmt"

type Animal interface {
	Says() string
	Counts() int
}

type Dog struct {
	Name  string
	Breed string
}

type Cat struct {
	Name  string
	Color string
	Age   int
}

func main() {

	var dog = Dog{"Bob", "Giwawa"}

	PrintInfo(&dog)

	var cat = Cat{"John", "Black", 5}

	PrintInfo(&cat)

}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and count is", a.Counts())
}

func (d *Dog) Says() string {
	return "wong!"
}

func (d *Dog) Counts() int {
	return 4
}

func (cat *Cat) Says() string {
	return cat.Color
}

func (cat *Cat) Counts() int {
	return cat.Age
}
