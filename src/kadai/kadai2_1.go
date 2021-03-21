// https://docs.microsoft.com/ja-jp/learn/modules/go-control-flow/5-challenge

package main

import (
	"fmt"
)

var Fizz int = 3
var Buzz int = 5
var Num int = 100

func main() {

	for i := 1; i <= Num; i++ {

		switch {
		case i%Fizz == 0 && i%Buzz == 0:
			fmt.Println("FizzBuzz")
		case i%Fizz == 0:
			fmt.Println("Fizz")
		case i%Buzz == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

	}
}
