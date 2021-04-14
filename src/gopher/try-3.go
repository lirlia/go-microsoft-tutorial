/*
* 素数判定(愚直)
 */

package main

import "fmt"

const (
	COUNT = 100
)

var sosu []int

func main() {

	for v := 2; v <= COUNT; v++ {

		// 2は素数なので初めから除外
		if v == 2 {
			fmt.Println(v)
			sosu = append(sosu, v)
			continue
		}

		// 既存の素数で割り切れないなら素数とみなす
		for i, r := range sosu {
			if v%r == 0 {
				break
			}

			if i == len(sosu)-1 {
				fmt.Println(v)
				sosu = append(sosu, v)
			}
		}
	}

}
