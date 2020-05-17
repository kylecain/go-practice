/*
Write a progam to check chether a given number is an ugly number.
Ugly numbers are POSITIVE NUMBERS  whose prime factors only include 2, 3, 5

Example 1:
Input: 6
Output: true
Explanation: 6 = 2 × 3

Example 2:
Input: 8
Output: true
Explanation: 8 = 2 × 2 × 2

Example 2:
Input: 8
Output: true
Explanation: 8 = 2 × 2 × 2

Note:
1 is typically treated as an ugly number.
Input is within the 32-bit signed integer range: [−231,  231 − 1]
*/

package main

func main() {

}

func isUgly(num int) bool {
	if num == 1 || num == 2 || num == 3 || num == 5 {
		return true
	}

	if num%2 != 0 && num%3 != 0 && num%5 != 0 {
		return false
	}

	if num%2 == 0 {
		return isUgly(num / 2)
	} else if num%3 == 0 {
		return isUgly(num / 3)
	} else if num%5 == 0 {
		return isUgly(num / 5)
	}

	return false
}
