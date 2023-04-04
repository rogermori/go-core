package main

import "fmt"

type Point struct {
	x, y int
}

func NewPoint(x,y int) Point {
	return Point{x,y}
}

type rect struct {
	pos Point
	width int
	height int
}

func (r rect) area() int {
	return r.height * r.width
}


func main() {
	p := Point{20,40}
	fmt.Println(p)
	
	r := rect {
		pos: NewPoint(20,40),
		width: 400,
		height: 300,
	}
	fmt.Println(r)
	fmt.Println(r.area())
	
}