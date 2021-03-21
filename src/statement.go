package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//	x := 27
	//	if x%2 == 0 {
	//		fmt.Println(x, "is even")
	//	}

	if num := givemenumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//---

	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)
	i = 0
	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("no match")
	}

	switch {
	case i == 0:
		fmt.Println("zero")
		fallthrough
	case i == 0:
		fmt.Println("zero2")
		fallthrough
	case i == 1:
		fmt.Println("one")
	case i == 2:
		fmt.Println("two")
	default:
		fmt.Println("no match")
	}

	fmt.Println("ok")
}

func givemenumber() int {
	return -7
}
