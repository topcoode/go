// Golang scope rules of variables can be divided into two categories which depend on where the variables are declared
// Local Variables(Declared Inside a block or a function)
// Global Variables(Declared outside a block or a function)
// LOCAL VARIABLE
// Variables that are declared inside a function or a block are termed as Local variables. These are not accessible outside the function or block.
// These variables can also be declared inside the for, while statement etc. inside a function.
// However, these variables can be accessed by the nested code blocks inside a function.
// These variables are also termed as the block variables.
// There will be a compile-time error if these variables are declared twice with the same name in the same scope.
// These variables don’t exist after the function’s execution is over.
// The variable which is declared outside the loop is also accessible within the nested loops. It means a global variable will be accessible to the methods and all loops. The local variable will be accessible to loop and function inside that function.
// A variable which is declared inside a loop body will not be visible to the outside of loop body.
package main

import "fmt"

// main function
func main() { // from here local level scope of main function starts

	// local variables inside the main function
	var myvariable1, myvariable2 int = 59, 488

	// Display the values of the variables
	fmt.Printf("The value of myvariable1 is : %d\n",
		myvariable1)

	fmt.Printf("The value of myvariable2 is : %d\n",
		myvariable2)

} // here local level scope of main function ends

/* OUTPUT:
The value of myvariable1 is : 89
The value of myvariable2 is : 45 */
// GLOBAL VARIABLES
/*
The variables which are defined outside of a function or a block is termed as Global variables.
These are available throughout the lifetime of a program.
These are declared at the top of the program outside all of the functions or blocks.
These can be accessed from any portion of the program.
*/
// global variable declaration
var myvariable1 int = 100

func globalvariable() { // from here local level scope starts

	// local variables inside the main function
	var myvariable2 int = 200

	// Display the value of global variable
	fmt.Printf("The value of Global myvariable1 is : %d\n",
		myvariable1)

	// Display the value of local variable
	fmt.Printf("The value of Local myvariable2 is : %d\n",
		myvariable2)

	// calling the function
	display()

} // here local level scope ends

// taking a function
func display() { // local level starts

	// Display the value of global variable
	fmt.Printf("The value of Global myvariable1 is : %d\n",
		myvariable1)

} // local scope ends here
//OUTPUT:
/* The value of Global myvariable1 is : 100
The value of Local myvariable2 is : 200
The value of Global myvariable1 is : 100 */
