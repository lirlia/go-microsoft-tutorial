/*
* switchを使ったおみくじ
 */

package main

import "fmt"

var (
	OMIKUJI = []string{"無効な値です", "凶", "吉", "吉", "中吉", "中吉", "大吉"}
	DICE    = 6
)

func main() {

	var result int
	for v := 0; v <= 10; v++ {

		// 標準入力を受け取る
		fmt.Scan(&result)

		// 出目によって出力変更
		switch result {
		case 1:
			fmt.Println(OMIKUJI[result])
		case 2, 3:
			fmt.Println(OMIKUJI[result])
		case 4, 5:
			fmt.Println(OMIKUJI[result])
		case 6:
			fmt.Println(OMIKUJI[result])
		default:
			fmt.Println(OMIKUJI[0])
		}
	}
}
