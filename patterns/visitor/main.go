package main

import (
	"fmt"
	"visitor/pkg"
)

func main() {
	shapes := []pkg.Shape{
		&pkg.Square{Side: 4},
		&pkg.Circle{Radius: 3},
		&pkg.Rectangle{Length: 2, Width: 3},
	}

	calculator := &pkg.AreaCalculator{}

	for _, shape := range shapes {
		shape.Accept(calculator)
	}
}

// Определение структуры и метода для Circle
type Circle struct {
	Radius int
}

func (c *Circle) Accept(v pkg.Visitor) {
	v.VisitForCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

// Определение структуры и метода для Square
type Square struct {
	Side int
}

func (s *Square) Accept(v pkg.Visitor) {
	v.VisitForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}

// Определение структуры и метода для Rectangle
type Rectangle struct {
	Length int
	Width  int
}

func (r *Rectangle) Accept(v pkg.Visitor) {
	v.VisitForRectangle(r)
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}

// Определение структуры для AreaCalculator
type AreaCalculator struct {
	Area int
}

func (a *AreaCalculator) VisitForSquare(s *pkg.Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) VisitForCircle(c *pkg.Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) VisitForRectangle(r *pkg.Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
