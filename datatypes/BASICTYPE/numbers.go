//NUMBERS are divided into three sub-categories
//>> Integers
//>> Floating-Point Numbers
// >> Complex Numbers
// INTERGERS::::------------->>>
//  In Go language, both signed and unsigned integers are available in four different sizes
// The signed int is represented by int and the unsigned integer is represented by uint.
/* int8	8-bit signed integer(-128 to 127)
int16	16-bit signed integer(-32,768 to 32,767)
int32	32-bit signed integer(-2,147,483,648 to 2,147,483,647)
int64	64-bit signed integer(-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)
uint8	8-bit unsigned integer(0 through 255 decimal)
uint16	16-bit unsigned integer(0 to 65535)
uint32	32-bit unsigned integer(2,147,483,647)
uint64	64-bit unsigned integer
int	Both int and uint contain same size, either 32 or 6-128 to 1274 bit.
uint	Both int and uint contain same size, either 32 or 64 bit.
rune	It is a synonym of int32 and also represent Unicode code points.
byte	It is a synonym of uint8.
uintptr	It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value. */
package main

import "fmt"

func main() {
	floatingno()
	// complex()
	var Numbers int = 77
	fmt.Println(Numbers)
	var Numbers1 int8 = 66 //128
	fmt.Println(Numbers1)
	var Numbers2 int16 = 3277
	fmt.Println(Numbers2)
	// ..............................>>>same ,they will maintain range.
}

// FLOATING_POINT NUMBERS
// In Go language, floating-point numbers are divided into two categories
/* float32	32-bit IEEE 754 floating-point number
float64	64-bit IEEE 754 floating-point number */
func floatingno() {
	a := 20.45
	b := 34.89

	// Subtraction of two
	// floating-point number
	c := b - a

	// Display the result
	fmt.Printf("Result is: %f", c)

	// Display the type of c variable
	fmt.Printf("\nThe type of c is : %T", c)
}

// COMPLEX NUMBERS:
// The complex numbers are divided into two parts
// float32 and float64 are also part of these complex numbers.
// in-built function creates a complex number from its imaginary and real part and in-built imaginary and real function extract those parts.
/* complex64	Complex numbers which contain float32 as a real and imaginary component.
complex128	Complex numbers which contain float64 as a real and imaginary component. */
// func complex() {

//    var a complex128 = complex(6, 2)
//    var b complex64 = complex(9, 2)
//    fmt.Println(a)
//    fmt.Println(b)

//    // Display the type
//   fmt.Printf("The type of a is %T and "+
//             "the type of b is %T", a, b)
// }
