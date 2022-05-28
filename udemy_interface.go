package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Triangle struct {
	base   float64
	height float64
	name   string `default:"Triangle"`
}

type Square struct {
	side float64
}

func (t Triangle) getArea() float64 {
	area := (t.base * t.height) / 2
	return area
}

func (s Square) getArea() float64 {
	area := s.side * s.side
	return area
}

func printArea(s Shape) {
	fmt.Println("Shape area is:", s.getArea())
}

func main() {
	triangle := Triangle{
		base:   10,
		height: 5,
	}
	square := Square{
		side: 7,
	}

	printArea(triangle)
	printArea(square)

}
