package main

import (
	"fmt"
	Helper "use_package/helpers"
)

func main() {
	// var myVar = Helper.SomeType{TypeName: "John", TypeNumber: 18}
	// var myVar = Helper.Says()
	// fmt.Println(myVar)
	var myVar Helper.SomeType
	myVar.TypeName = "some name"
	fmt.Println(myVar.TypeName)

}
