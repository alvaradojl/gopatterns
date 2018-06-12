package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("starting...")

	num3 := 3
	doubleNum := func() int {
		//this is known as closure
		num3 *= 2
		return num3
	}

	rect1 := Rectangle{height: 100, width: 50}
	circ := Circle{4}

	fmt.Printf("\n shape area of rectangle: %f", getArea(rect1))
	fmt.Printf("\n shape area of circle: %f", getArea(circ))

	fmt.Printf("\n area: %f", rect1.width)

	fmt.Printf("\narea is %f: \n", rect1.area())

	fmt.Printf("\ndoubleNum result: %d\n", doubleNum())

}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}
