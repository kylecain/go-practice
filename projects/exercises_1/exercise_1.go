package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
}

//Using the short declaration operator, ASSIGN these VALUES to
//VARIABLES with the IDENTIFIERS "x" and "y" and "z"
func exercise1() {
	x := 42
	y := "James Bond"
	z := true

	//One print statement
	fmt.Printf("%v %v %v", x, y, z)

	//Multiple print statements
	fmt.Println("\n", x)
	fmt.Println(y)
	fmt.Println(z)
}

//Use var to DECLARE three VARIABLES. The variables should have package level scope.
//Do not assign VALUES to the variables. Use the
//following IDENTIFIERS for the variables and make sure the variables
//are of the following TYPE (meaning they can store VALUES of that TYPE).

var x int
var y string
var z bool

func exercise2() {
	//print out the values for each identifier
	//The compiler assigned values to the variables. What are these values called?
	fmt.Printf("%v %v %v", x, y, z)

	//the default values assigned are called the Zero Values
}

/*
Using the code from the previous exercise,
At the package level scope, assign the following values to the three variables
for x assign 42
for y assign “James Bond”
for z assign true
in func main
use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
print out the value stored by variable “s”
*/

var xx int = 42
var yy string = "James Bond"
var zz bool = true

func exercise3() {
	s := fmt.Sprintf("\n%v %v %v", xx, yy, zz)
	fmt.Println(s)
}

/*
Create your own type. Have the underlying type be an int.
create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
in func main
print out the value of the variable “x”
print out the type of the variable “x”
assign 42 to the VARIABLE “x” using the “=” OPERATOR
print out the value of the variable “x”
*/
type myType int

func exercise4() {
	var xxx myType
	fmt.Printf("\n%v %T", xxx, xxx)
	xxx = 42
	fmt.Printf("\n%v\n", xxx)
}

/*
Building on the code from the previous example
at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”.
The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
in func main
this should already be done
print out the value of the variable “x”
print out the type of the variable “x”
assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
print out the value of the variable “x”
now do this
now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
then use the “=” operator to ASSIGN that value to “y”
print out the value stored in “y”
print out the type of “y”
*/
var xxxx myType
var yyyy int

func exercise5() {
	fmt.Printf("%v %T", xxxx, xxxx)
	xxxx = 10
	fmt.Printf("\n%v\n", xxxx)

	yyyy = int(xxxx)
	fmt.Printf("\n%v %T\n", yyyy, yyyy)
}
