package main

import (
	"os"
	"strconv"
)

func main() {

	//number1, _ := strconv.Atoi(os.Args[1])
	//number2, _ := strconv.Atoi(os.Args[2])
	//println("sum:", number1+number2)

	println(sum(os.Args[1], os.Args[2]))
	println(calc(os.Args[1], os.Args[2]))
}

func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}

func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}
