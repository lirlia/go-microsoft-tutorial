package main

// 数を尋ねて、負の場合はパニックにする

import (
	"fmt"
)

func main() {
	val := 0

	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("you entered negative number !!!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}
