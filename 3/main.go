// types and interfaces, and unit testing

package main

import (
	"fmt"
	"math"
)

// Square Type has Width and Height as parameters
type Square struct {
	Width float64
	Height float64
}

// Circle Type has Radius as parameter
type Circle struct {
	Radius float64
}

// Shape interface has area() method which calculates are of shape depending wether its Square or Circle
type shape interface {
	area() float64
}

func (s Square) area() float64 {
	return s.Width * s.Height
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.Radius,2)
}

// Info method returns calculated area of the input shape
func Info(s shape) float64 {
	return s.area()
}

func main() {

	square1 := Square{
		Width: 12,
		Height: 10,
	}

	circle1 := Circle {
		Radius: 5,
	}

	fmt.Println(Info(square1))
	fmt.Println(Info(circle1))

}

