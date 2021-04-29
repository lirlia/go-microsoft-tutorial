package main

type MyInt int

func (n *MyInt) Inc() {
	println("show:", n)
	*n++
}
func (n MyInt) Inc2() {
	println("show:", n)
	n++
}
func main() {
	var n MyInt
	println(n)
	n.Inc()
	n.Inc2()
	println(n)
}
