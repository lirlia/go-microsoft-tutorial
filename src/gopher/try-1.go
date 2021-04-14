/*
* ifを使った奇数/偶数判定
 */

package main

import "fmt"

const (
	ODD   = "奇数"
	EVEN  = "偶数"
	COUNT = 100
)

func main() {

	var result string
	for v := 0; v <= COUNT; v++ {

		// 偶数or奇数で出力結果を変更する
		if v%2 == 0 {
			result = EVEN
		} else {
			result = ODD
		}

		fmt.Printf("%v-%v\n", v, result)

	}

}
