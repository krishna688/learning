// Question 3: Interfaces
// Define an interface Shape with two methods: Area() and Perimeter(). Then, create two structs, Rectangle and Circle, and implement the Shape interface for both.

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	l, b int
}

func (r *Rectangle) Area() {
	fmt.Println(r.l * r.b)
}
func (r *Rectangle) Perimeter() {
	fmt.Println(2 * (r.l + r.b))
}

type Circle struct {
	s float64
}

func (c *Circle) Area() {
	fmt.Println(math.Pi * c.s * c.s)
}
func (c *Circle) Perimeter() {
	fmt.Println(2 * math.Pi * c.s)
}

func main() {
	var k Shape
	k = &Rectangle{2, 3}
	k.Area()
	k.Perimeter()
	k = &Circle{5}
	k.Area()
	k.Perimeter()
}
