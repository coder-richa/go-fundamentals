package main

import (
	"math"
)

// Interface Declaration
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r *Rectangle) SetLength(length float64) {
	r.length = length
}

func (r *Rectangle) SetBreadth(breadth float64) {
	r.breadth = breadth
}

func (r *Rectangle) Length() float64 {
	return r.length
}

func (r *Rectangle) Breadth() float64 {
	return r.breadth
}

/* Attaching Shape Interface Methods */
func (r *Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

type Circle struct {
	radius float64
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

func (c *Circle) Radius() float64 {
	return c.radius
}

/* Attaching Shape Interface Methods */
func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return math.Pi * 2 * c.radius
}
