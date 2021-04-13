package main

import (
	"fmt"
	"math"
	//"github.com/myuser/geometory"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Triangle struct {
	size int
}

type Colored_triangle struct {
	Triangle
	color string
}

type Square struct {
	size float64
}

func (t Triangle) Perimeter() int {
	return t.size * 3
}

func (t *Triangle) SetSize(size int) {
	t.size = size
}

func (t *Triangle) DoubleSize() {
	t.size *= 2
}

func (t Colored_triangle) Perimeter() int {
	return t.size * 30
}

func (t Square) Perimeter() float64 {
	return t.size * 4
}

func (t Square) Area() float64 {
	return t.size * 2
}

func (t *Colored_triangle) SetColor(color string) {
	t.color = color
}

///
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printInfomation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area", s.Area())
	fmt.Println("Perimeter", s.Perimeter())
	fmt.Println()
}

///

func main() {
	t := Triangle{}
	t.SetSize(3)
	//s := Square{3}
	c := Colored_triangle{}
	c.SetSize(3)
	c.SetColor("blue")

	fmt.Println("colored_triangle:", c.Perimeter())

	fmt.Println("triangle:", t.Perimeter())
	t.DoubleSize()
	fmt.Println("triangle:", t.Perimeter())
	//fmt.Println("square:", s.perimeter())

	var s Shape = Square{3}
	printInfomation(s)

	cc := Circle{6}
	printInfomation(cc)
}
