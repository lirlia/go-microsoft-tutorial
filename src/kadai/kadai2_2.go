// https://docs.microsoft.com/ja-jp/learn/modules/go-control-flow/5-challenge
// [平方根を推測する]
//    数の平方根を推測するプログラムを作成します。 次の数式を使用します。
//    sroot = sroot − (sroot − x) / (2 * sroot)

package main

import (
	"fmt"
	"os"
	"strconv"
)

var LOOP_COUNT int = 10

func calcSroot(n float64, sroot float64) float64 {
	return sroot - (sroot*sroot-n)/(2*sroot)
}

func main() {

	num, _ := strconv.ParseFloat(os.Args[1], 64)
	currguess := 1.0
	prevguess := 0.0

	for i := 0; i < LOOP_COUNT; i++ {
		prevguess = currguess
		currguess = calcSroot(num, prevguess)

		if prevguess == currguess {
			break
		}
		fmt.Println("A guess for square root is", currguess)
	}
}
