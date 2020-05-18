package main

import "fmt"

func main() {
	//exercise1()
	//exercise2()
	//exercise3()
	//exercise4()
	//exercise5()
	//exercise6()
	//exercise7()
	//exercise8()
	exercise9()
	exercise10()

}

/*
Hands-on exercise #1
Using a COMPOSITE LITERAL:
create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position.
Range over the array and print the values out.
Using format printing
print out the TYPE of the array
*/

func exercise1() {
	var a = [5]int{2, 4, 6, 8, 10}
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", a)
}

/*
Hands-on exercise #2
Using a COMPOSITE LITERAL:
create a SLICE of TYPE int
assign 10 VALUES
Range over the slice and print the values out.
Using format printing
print out the TYPE of the slice
*/

func exercise2() {
	var s = []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", s)
}

/*
Hands-on exercise #3
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]

*/

func exercise3() {
	var s = []int{42, 43, 44, 45, 46}
	var s1 = []int{47, 48, 49, 50, 51}
	var s2 = []int{44, 45, 46, 47, 48}
	var s3 = []int{43, 44, 45, 46, 47}

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

/*
Hands-on exercise #4
Follow these steps:
start with this slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
append to that slice this value
52
print out the slice
in ONE STATEMENT append to that slice these values
53
54
55
print out the slice
append to the slice this slice
y := []int{56, 57, 58, 59, 60}
print out the slice
*/

func exercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

/*
Hands-on exercise #5
To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
start with this slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
[42, 43, 44, 48, 49, 50, 51]

*/

func exercise5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[3+3:]...)
	fmt.Println(y)
}

/*
Hands-on exercise #6
Create a slice to store the names of all of the states in the United States of America. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position in the slice, without using the range clause. Here is a list of the states:

` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,

*/

func exercise6() {
	s := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println("There are this many states: ", len(s))
	fmt.Println("Capacity of slice: ", cap(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("index: %d\t\tstate: %v\n", i, s[i])
	}
}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.
*/

func exercise7() {
	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	s := [][]string{s1, s2}
	for _, v := range s {
		for _, w := range v {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}

/*
Hands-on exercise #8
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

    `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

*/

var m = map[string][]string{
	`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
	`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
	`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
}

func exercise8() {

	for k, v := range m {
		for i, w := range v {
			fmt.Printf("%v likes %v\tindex position:%d\n", k, w, i)
		}
	}
}

/*
Hands-on exercise #9
Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

*/

func exercise9() {
	m["kyle"] = []string{`headphones`, `league of legends`, `memes`}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
Hands-on exercise #10
Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
*/

func exercise10() {
	delete(m, "kyle")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
