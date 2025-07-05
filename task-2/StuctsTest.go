package main

import (
	"fmt"
	"math"
)

func main() {
	shape := Rectangle{width: 5, height: 10}
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
	circle := Circle{radius: 5}
	fmt.Println(circle.Area())
	fmt.Println(circle.Perimeter())
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
