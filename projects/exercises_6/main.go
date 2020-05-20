package main

import "fmt"

func main() {
	//exercise 3
	defer deferFunc()

	//exercise 1
	int1 := foo()
	int2, str := bar()

	fmt.Printf("int1 is equal to %v\nint2 is equal to %v\nstr is equal to %v\n", int1, int2, str)

	//exercise 2
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(foo1(xi...))
	fmt.Println(bar1(xi))

	p1 := person{
		first: "bob",
		last:  "bobbert",
		age:   43,
	}
	p1.speak()

	myCircle := circle{
		radius: 400.0,
	}

	mySquare := square{
		length: 398.0,
		width:  201,
	}
	fmt.Print(info(myCircle))
	fmt.Print(info(mySquare))

	f := func() string {
		return "exercise 6 and 7 i guess\n"
	}
	fmt.Print(f())

	f1 := firstFunc()
	f1()

	fmt.Print(callbackFunc(f))
}

/*
Hands on exercise
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results

*/

func foo() int {
	return 1
}
func bar() (int, string) {
	return 1, "hi"
}

/*
Hands-on exercise #2
create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in

create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in
*/

func foo1(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar1(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

/*
exercise 3

Use the “defer” keyword to show that a deferred func runs after the func containing it exits
*/
func deferFunc() {
	fmt.Println("i was defered! and part of exercise 3")
}

/*
Hands-on exercise #4
Create a user defined struct with
the identifier “person”
the fields:
first
last
age
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person
*/

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi im %v %v and i am %v years old\n", p.first, p.last, p.age)
}

/*
Hands-on exercise #5
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/

type circle struct {
	radius float32
}

type square struct {
	length float32
	width  float32
}

func (c circle) area() float32 {
	return c.radius * 2 * 3.14
}

func (s square) area() float32 {
	return s.length * s.width
}

type shape interface {
	area() float32
}

func info(s shape) string {
	return fmt.Sprintln(s.area())
}

/*
Hands-on exercise #8
Create a func which returns a func
assign the returned func to a variable
call the returned func
*/
func firstFunc() func() {
	return func() {
		fmt.Println("do the thing")
	}
}

/*
Hands-on exercise #9
A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument
*/

func callbackFunc(f func() string) string {
	return fmt.Sprintln(f(), "part of callback")
}

/*
Hands-on exercise #10
Closure is when we have “enclosed” the scope of a variable in some code block.
For this hands-on exercise, create a func which “encloses” the scope of a variable:
*/
//they all do
