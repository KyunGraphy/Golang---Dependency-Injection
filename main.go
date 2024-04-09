package main

import (
	"fmt"
	"math"
)

// Initialize the structure of shapes
type Circle struct {
	radius float64
	calc   Calc
}

type Rectangle struct {
	width  float64
	height float64
	calc   Calc
}

// Initialize the methods interface of shapes
type Calc interface {
	perimeter() float64
	area() float64
}

// Initial new shape method
func NewCircle(r float64) *Circle {
	return &Circle{
		radius: r,
	}
}

func NewRectangle(w float64, h float64) *Rectangle {
	return &Rectangle{
		width:  w,
		height: h,
	}
}

// Define the logic of shape methods
func (shape Circle) perimeter() float64 {
	return 2 * math.Pi * shape.radius
}

func (shape Circle) area() float64 {
	return math.Pi * math.Pow(shape.radius, 2)
}

func (shape Rectangle) perimeter() float64 {
	return 2 * (shape.height + shape.width)
}

func (shape Rectangle) area() float64 {
	return shape.height * shape.width
}

// Main goroutine
func main() {
	circle := NewCircle(4)
	fmt.Println(circle.perimeter())
	fmt.Println(circle.area())

	rectangle := NewRectangle(6, 8)
	fmt.Println(rectangle.perimeter())
	fmt.Println(rectangle.area())
}
