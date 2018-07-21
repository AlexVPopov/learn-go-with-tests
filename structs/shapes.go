package main

import "math"

func main() {}

// Shape is an interface for shapes
type Shape interface {
	Area() float64
}

// Rectangle represents the geometrical shape, known as rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle
type Circle struct {
	Radius float64
}

// Area returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle represents a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}
