// The boolean data type represents only one bit of information either true or false.
// The values of type boolean are not converted implicitly or explicitly to any other type.
package main

import "fmt"

func main() {

	// variables
	str1 := "apple"
	str2 := "orange"
	str3 := "banana"
	result1 := str1 == str2
	result2 := str1 == str3

	// Display the result
	fmt.Println(result1)
	fmt.Println(result2)

	// Display the type of
	// result1 and result2
	fmt.Printf("The type of result1 is %T and "+
		"the type of result2 is %T",
		result1, result2)

}

/* false
false
The type of result1 is bool and the type of result2 is bool */
