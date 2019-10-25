// Package circleutil with functions for circles
package circleutil

import "math"

// Circle with center (x,y) and radius r
type Circle struct {
	x int
	y int // x,y center of circle
	r float64
}

// Area calculates the area of circle
func (cir Circle) Area() float64 {
	return math.Pi * math.Pow(cir.r, 2)
}

// NewCir makes new Circle with center (x,y) and radius r
func NewCir(x, y int, r float64) Circle {
	return Circle{x: x, y: y, r: r}
}
