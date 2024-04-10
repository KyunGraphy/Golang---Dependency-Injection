package main

import (
	"fmt"
	"math"
)

// 1.Initialize the structure of shapes
type Circle struct {
	radius float64
	calc   Calc
}

type Rectangle struct {
	width  float64
	height float64
	calc   Calc
}

// 2.Initialize the methods interface of shapes
type Calc interface {
	perimeter() float64
	area() float64
}

// 3.Create specific implementation
type CircleImpl struct {
	radius float64
}
type RectangleImpl struct {
	width, height float64
}

// 4.Initial new shape method
func NewCircle(r float64, c Calc) *Circle {
	return &Circle{
		radius: r,
		calc:   c,
	}
}

func NewRectangle(w float64, h float64, c Calc) *Rectangle {
	return &Rectangle{
		width:  w,
		height: h,
		calc:   c,
	}
}

// 5.Define the logic of shape methods
func (s *CircleImpl) perimeter() float64 {
	return 2 * math.Pi * (s.radius)
}

func (s *CircleImpl) area() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}

func (r *RectangleImpl) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r *RectangleImpl) area() float64 {
	return r.height + r.width
}

// 6.Define consumer methods provider
func (c *Circle) perimeter() float64 {
	return c.calc.perimeter()
}

func (c *Circle) area() float64 {
	return c.calc.area()
}

func (r *Rectangle) perimeter() float64 {
	return r.calc.perimeter()
}

func (r *Rectangle) area() float64 {
	return r.calc.area()
}

// Main goroutine
func main() {
	circleConsumer := NewCircle(4, &CircleImpl{radius: 4})
	fmt.Println(circleConsumer.perimeter())
	fmt.Println(circleConsumer.area())

	rectangle := NewRectangle(6, 8, &RectangleImpl{width: 6, height: 8})
	fmt.Println(rectangle.perimeter())
	fmt.Println(rectangle.area())
}
