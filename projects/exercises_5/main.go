package main

import "fmt"

func main() {
	//exercise1()
	//exercise2()
	//exercise3()
	exercise4()

}

/*
Hands-on exercise #1
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/

func exercise1() {
	type person struct {
		fName       string
		lName       string
		favIceCream []string
	}
	p1 := person{
		fName:       "bob",
		lName:       "bobert",
		favIceCream: []string{"vanilla", "chocolate"},
	}
	p2 := person{
		fName:       "wendy",
		lName:       "bobert",
		favIceCream: []string{"mint", "buttercream"},
	}

	xp := []person{p1, p2}

	for _, v := range xp {
		fmt.Println(v.fName)
		fmt.Println(v.lName)
		for _, v := range v.favIceCream {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

/*
Hands-on exercise #2
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
*/

func exercise2() {
	type person struct {
		fName       string
		lName       string
		favIceCream []string
	}
	p1 := person{
		fName:       "bob",
		lName:       "bobert",
		favIceCream: []string{"vanilla", "chocolate"},
	}
	p2 := person{
		fName:       "wendy",
		lName:       "sobert",
		favIceCream: []string{"mint", "buttercream"},
	}
	xp := []person{p1, p2}
	//var m map[string]person
	m := make(map[string]person)

	for _, v := range xp {
		m[v.lName] = v
	}
	for _, v := range m {
		fmt.Println(v.fName)
		fmt.Println(v.lName)
		for _, v := range v.favIceCream {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

/*
Hands-on exercise #3
Create a new type: vehicle.
The underlying type is a struct.
The fields:
doors
color
Create two new types: truck & sedan.
The underlying type of each of these new types is a struct.
Embed the “vehicle” type in both truck & sedan.
Give truck the field “fourWheel” which will be set to bool.
Give sedan the field “luxury” which will be set to bool. solution
Using the vehicle, truck, and sedan structs:
using a composite literal, create a value of type truck and assign values to the fields;
using a composite literal, create a value of type sedan and assign values to the fields.
Print out each of these values.
Print out a single field from each of these values.

*/

func exercise3() {
	type vehicle struct {
		doors int
		color string
	}
	type truck struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "silver",
		},
		luxury: false,
	}
	fmt.Println(t1)
	fmt.Println(t1.vehicle.color)
	fmt.Println(s1)
	fmt.Println(t1.vehicle.doors)

}

/*
Hands-on exercise #4
Create and use an anonymous struct.
*/

func exercise4() {
	keyboard := struct {
		keycaps  string
		switches string
		size     string
	}{
		keycaps:  "abs",
		switches: "box whites",
		size:     "tenkeyless",
	}
	fmt.Println(keyboard)
	fmt.Println(keyboard.switches)

}
