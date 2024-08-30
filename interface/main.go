package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	area() float64
	perimeter() float64

	// printHello - if the next line is commented back in we will get an error when calling 'Measure(shape1)'
	//printHello()
}

func Measure(shape shape) {
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	fmt.Printf("Type: %v\n", reflect.TypeOf(shape))
	fmt.Printf("Area: %f\n", shape.area())
	fmt.Printf("Perimeter: %f\n", shape.perimeter())
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
}

func main() {
	shape1 := circle{
		radius: 3,
	}

	// printHello - if 'printHello' is commented back in the interface we will get an error. NOTE: The method is present in 'rectangle' so no error there.
	// Cannot use 'shape1' (type circle) as the type Shape Type does not implement 'Shape' as some methods are missing: printHello()
	Measure(shape1)
	// In intellij we can just use option+Enter to 'Implement missing methods'.

	shape2 := rectangle{
		width:  2,
		height: 2,
	}

	// printHello - if 'printHello' is commented back in the interface we will not get an error here.
	Measure(shape2)

}
