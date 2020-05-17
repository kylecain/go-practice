package main

import "fmt"

func main() {
	num := 20
	exercise1(num)
	exercise2(num)
	exercise3()
	exercise4(num)
	exercise5()
	exercise6()
}

/*
Hands-on exercise #1
Write a program that prints a number in decimal, binary, and hex
*/
func exercise1(n int) {
	fmt.Printf("Decimal: %d\nBinary: %b\nHexidecimal: %x\n", n, n, n)
}

/*
Hands-on exercise #2
Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
Now print each of the variables.
*/

func exercise2(n int) {
	a := (n == n)
	b := (n <= n)
	c := (n >= n)
	d := (n != n)
	e := (n > n)
	f := (n < n)

	fmt.Printf("(%d == %d) is %t\n(%d <= %d) is %t\n(%d >= %d) is %t\n(%d != %d) is %t\n(%d > %d) is %t\n(%d < %d) is %t\n", n, n, a, n, n, b, n, n, c, n, n, d, n, n, e, n, n, f)
}

/*
Hands-on exercise #3
Create TYPED and UNTYPED constants. Print the values of the constants.
*/
const untypedConst = 40
const typedCost int = 50

func exercise3() {
	fmt.Println(untypedConst, typedCost)
}

/*
Hands-on exercise #4
Write a program that
assigns an int to a variable
prints that int in decimal, binary, and hex
shifts the bits of that int over 1 position to the left, and assigns that to a variable
prints that variable in decimal, binary, and hex
*/

func exercise4(n int) {
	myInt := n
	fmt.Printf("Decimal: %d\nBinary: %b\nHexidecimal: %x\n", myInt, myInt, myInt)
	myShiftedInt := n << 1
	fmt.Printf("Decimal: %d\nBinary: %b\nHexidecimal: %x\n", myShiftedInt, myShiftedInt, myShiftedInt)
}

/*
Hands-on exercise #5
Create a variable of type string using a raw string literal. Print it.
*/

func exercise5() {
	var myString string = `This is a "string literal"`
	fmt.Println(myString)
}

/*
Hands-on exercise #6
Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
*/
const (
	nextYear             = 2020 + iota
	nextNextYear         = 2020 + iota
	nextNextNextYear     = 2020 + iota
	nextNextNextNextYear = 2020 + iota
)

func exercise6() {
	fmt.Printf("Next year is: %d\nThe year after that is: %d\nThe year after that is: %d\nThe year after that is: %d\n", nextYear, nextNextYear, nextNextNextYear, nextNextNextNextYear)
}
