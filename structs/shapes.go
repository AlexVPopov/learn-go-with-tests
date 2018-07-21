package main

func main() {}

// Rectangle is a type, representing the geometrical shape, known as rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter takes the width and height of a rectangle and returns its perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area takes the width and height of a rectangle and returns its area
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
