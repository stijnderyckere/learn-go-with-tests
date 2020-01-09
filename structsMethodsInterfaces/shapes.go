package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

//Rectangle : contains properties of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle : contains properties of a circle
type Circle struct {
	Radius float64
}

//Shape : interface that accepts any struct with Area as function
type Shape interface {
	Area() float64
}

//Perimeter : gets the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area : gets the area of a passed rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

//Area : gets the area of a passed rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Area : gets the area of a passed circle
func (c Circle) Area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}
