package main

import (
	"fmt"
	"math"
)

// rectangle - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
type rectangle struct {
	width, height float64
}

func (rect rectangle) area() float64 {
	return rect.width * rect.height
}

func (rect rectangle) perimeter() float64 {
	return 2*rect.width + 2*rect.height
}

func (rect rectangle) printHello() {
	fmt.Println("hello")
}

// Circle - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
