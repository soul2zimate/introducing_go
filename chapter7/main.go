package main

import (
	"fmt"
	"math"
)

// Circle  struct
type Circle struct {
	x, y, r float64
}

// Rectangle struct
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleAreaPtr(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func modify(c *Circle) {
	c.r = 6
}

// method
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// Person embedded types
type Person struct {
	name string
}

func (p *Person) talk() {
	fmt.Println("hi my name is : ", p.name)
}

// Android struct
type Android struct {
	Person
	model string
}

// interface

// Shape interface
type Shape interface {
	area() float64
	perimeter() float64
}

// interface
func totleArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// MultiShape struct
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, v := range m.shapes {
		area += v.area()
	}
	return area
}

// exercise
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * l * w
}

func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {
	// var c Circle
	// c := new(Circle)
	// c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)
	c.x = 10
	fmt.Println(c.x, c.y, c.r)

	fmt.Println(circleArea(c))
	fmt.Println(circleAreaPtr(&c))

	modify(&c)
	fmt.Println(c.x, c.y, c.r)

	// method
	fmt.Println("c.area : ", c.area())

	// embedded types

	a := new(Android)
	a.talk()

	// interface
	r := Rectangle{0, 0, 10, 10}
	fmt.Println("total area: ", totleArea(&c, &r))

	multiShapes := MultiShape{
		shapes: []Shape{
			&Circle{0, 0, 5},
			&Rectangle{0, 0, 10, 10},
		},
	}
	fmt.Println("multiShapes area: ", multiShapes.area())

	//exercise
	fmt.Println("multiShapes perimeter: ", multiShapes.perimeter())
}
