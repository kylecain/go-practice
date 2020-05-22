package main

import "fmt"

func main() {
	exercise1()
	p := person{
		first: "bob",
		last:  "bobbert",
		age:   43,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

/*
Hands-on exercise #1
Create a value and assign it to a variable.
Print the address of that value.

*/
func exercise1() {
	v := 12
	fmt.Println(&v)
}

/*
Hands-on exercise #2
create a person struct
create a func called “changeMe” with a *person as a parameter
change the value stored at the *person address
important
to dereference a struct, use (*value).field
p1.first
(*p1).first
“As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
https://golang.org/ref/spec#Selectors
create a value of type person
print out the value
in func main
call “changeMe”
in func main
print out the value
*/

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "bobthony"
	p.last = "bobtano"
}
