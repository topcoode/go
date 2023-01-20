package main

import (
	"fmt"
	"reflect"
)

func main() {
	var NAME string = "hanuma" //using var keyword
	//declaration of variable //variable declaration using intial value
	name1 := "hanuman"    //using ":="keyword
	var secondname string //variable declaration without using intial value
	fmt.Println("", NAME, name1, secondname)
	fmt.Println(reflect.TypeOf(NAME))

	//----------------------------------------
	//types of variables(datatypes):A variable is a storage unit of a particular data type.
	// You can give assign a name (label) to this storage unit.
	// A variable must be defined with the type of data it is holding.
	// For example,
	// string is a data type provided by Go to store character or text data.
	//If you declare a variable without assigning it a value,
	// Golang will automatically bind a default value (or a zero-value) to the variable.
	var name3 int = 56 //integer
	fmt.Println(name3)
	var name4 bool = true // boolean
	fmt.Println(name4)
	var name5 float32 = 55.55
	fmt.Println(name5)
	//---------------------------------------
	//Declaration with type inference
	var name2 = 25
	fmt.Println(name2)
	//----------------------------------------
	//Declaration of multiple variables
	var name6, name7 string = "sai", "vennu"
	fmt.Println(name6 + name7)
	//----------------------------------------
	//var s = "Japan"
	//Scope of Golang Variables Defined by Brace Brackets

	x := true

	if x {
		y := 1
		if x != false {
			//fmt.Println(s)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
	//--------------------------------------
	//Golang Variable Declaration Block
	var (
		product  = "Mobile"
		quantity = 50
		price    = 50.50
		inStock  = true
	)
	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(product)
	fmt.Println(inStock)
}
