package main

func swap(x int, y int) (fx int, fy int) {
	fx, fy = y, x
	return
}

func main() {
	n, m := swap(10, 20)
	println(n, m)
}
