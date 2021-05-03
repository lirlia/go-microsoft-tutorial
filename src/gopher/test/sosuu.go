/*
* 素数判定(愚直)
 */

package main

import "fmt"

func sosuchk(n int) []int {
	var sosu []int
	if n < 2 {
		return sosu
	}
	for v := 2; v <= n; v++ {

		// 2は素数なので初めから除外
		if v == 2 {
			sosu = append(sosu, v)
			continue
		}

		// 既存の素数で割り切れないなら素数とみなす
		for i, r := range sosu {
			if v%r == 0 {
				break
			}

			if i == len(sosu)-1 {
				sosu = append(sosu, v)
			}
		}
	}
	return sosu
}
func main() {
	fmt.Println(sosuchk(10))
}
