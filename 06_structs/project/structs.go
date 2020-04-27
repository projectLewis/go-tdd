package main

import (
	"math"
)

// Shape is an interface for shapes with area
type Shape interface {
	Area() float64
}

// Rectangle has a width and height
type Rectangle struct {
	Width float64
	Height float64
}
// Perimeter calculate the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle has a radius
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle has a base and height
type Triangle struct {
    Base   float64
    Height float64
}

// Area calculates the area of a circle
func (t Triangle) Area() float64 {
	half := 0.5
	return t.Base * t.Height * half
}
