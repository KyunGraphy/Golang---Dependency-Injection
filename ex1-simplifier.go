package main

import (
    "fmt"
    "math"
)

// 1.Initialize the structure of shapes
type Shape struct {
    methods Method
}

// 2.Initialize the methods interface of shapes
type Method interface {
    perimeter() float64
    area() float64
}

// 3.Create specific implementation
type Circle struct {
    radius  float64
}

type Rectangle struct {
    width   float64
    height  float64
}

// 4.Initial new shape method
func NewCircle (method Method) *Shape {
    return &Shape{
        methods: method,
    }
}

func NewRectangle (method Method) *Shape {
    return &Shape{
        methods: method,
    }
}

// 5.Define the logic of shape methods
func (c *Circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func (r *Rectangle) area() float64 {
    return r.width * r.height
}

func (r *Rectangle) perimeter() float64 {
    return 2 * (r.width + r.height)
}

// 6.Define consumer methods provider
func (s *Shape) area() float64 {
    return s.methods.area()
}

func (s *Shape) perimeter() float64 {
    return s.methods.perimeter()
}

// Main goroutine
func main() {
    rectangle := NewRectangle(&Rectangle{3, 4})
    fmt.Println(rectangle.area())
    fmt.Println(rectangle.perimeter())
    circle := NewCircle(&Circle{5})
    fmt.Println(circle.area())
    fmt.Println(circle.perimeter())
}
