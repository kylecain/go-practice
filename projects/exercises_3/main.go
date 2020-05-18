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
	//exercise10()

}

/*
Hands-on exercise #1
Print every number from 1 to 10,000
*/

func exercise1() {
	for i := 1; i <= 1000; i++ {
		fmt.Println(i)
	}
}

/*
Hands-on exercise #2
Print every rune code point of the uppercase alphabet three times. Your output should look like this:
65
    U+0041 'A'
    U+0041 'A'
U+0041 'A'
66
    U+0042 'B'
    U+0042 'B'
    U+0042 'B'
 … through the rest of the alphabet characters

*/

func exercise2() {
	for i := 65; i <= 90; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
}

/*
Hands-on exercise #3
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.

*/

func exercise3() {
	i := 1997
	for i <= 2020 {
		fmt.Println(i)
		i++
	}
}

/*
Hands-on exercise #4
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
*/

func exercise4() {
	i := 1997
	for {
		if i > 2020 {
			break
		}
		fmt.Println(i)
		i++
	}
}

/*
Hands-on exercise #5
Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
*/

func exercise5() {

	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

/*
Hands-on exercise #6
Create a program that shows the “if statement” in action.
*/

func exercise6() {

	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}

/*
Hands-on exercise #7
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
*/

func exercise7() {

	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		} else if i%5 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("hello")
		}
	}
}

/*
Hands-on exercise #8
Create a program that uses a switch statement with no switch expression specified.
*/

func exercise8() {
	a, b := 1, 2
	switch {
	case a == b:
		fmt.Println("a = b")
	case a != b:
		fmt.Println("a != b")
	}
}

/*
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
*/

func exercise9() {
	favSport := "basketball"
	switch favSport {
	case "volleyball":
		fmt.Println("me too thanks")
	case "basketball":
		fmt.Println("sad")
	}
}

/*
Hands-on exercise #10
Write down what these print:
fmt.Println(true && true)		true
fmt.Println(true && false) 		false
fmt.Println(true || true) 		true
fmt.Println(true || false) 		true
fmt.Println(!true)				false
*/

func exercise10() {
}
