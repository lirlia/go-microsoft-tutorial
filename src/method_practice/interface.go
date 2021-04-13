package main

import (
	"fmt"
	"strconv"
)

//構造体の定義
type MyCar struct {
	name  string
	speed int
}

//メソッドの実装
func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "kmで走ります"
}
func (u *MyCar) stop() {
	fmt.Println("停止します")
	u.speed = 0
}

func (u *MyCar) show() {
	fmt.Println("現在のスピード:" + strconv.Itoa(u.speed))
}

// ポインタを受け取らない形にメソッドを作成する
// これにより値渡しで処理をされるのでここで変更した内容は
// インスタンスに影響しない
func (u MyCar) run_test(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "kmで走ります"
}

func main() {
	c := &MyCar{name: "aaa", speed: 0}
	fmt.Println(c.run(50))
	c.show()
	fmt.Println(c.run_test(30))
	c.show()
	fmt.Println("↑30kmは値渡しで実現しているので、インスタンス変数は変更されない")
	c.stop()
	c.show()
}
