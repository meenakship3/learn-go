package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return math.Pi * 2 * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f\n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f\n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{15.2},
		circle{7.5},
		square{4.9},
		circle{12.3},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
