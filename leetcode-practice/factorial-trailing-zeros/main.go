package main

import (
	"fmt"
	"math/big"
)

func main() {

}
func TrailingZeroes(n int) int {
	num := big.NewInt(1)
	for i := n; i > 0; i-- {
		k := big.NewInt(int64(i))
		num.Mul(k, num)
	}

	strNum := fmt.Sprintf("%v", num)

	var c int
	for j := len(strNum) - 1; j >= 0; j-- {
		m := strNum[j] - 48
		if m != 0 {
			return c
		}
		if m == 0 {
			c++
		}
	}
	return c
}
