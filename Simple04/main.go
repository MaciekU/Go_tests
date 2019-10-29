package main

import (
	"fmt"
	"math"

	"github.com/maciek13U/Go_tests/Simple01/circleutil"
)

type Circle struct {
	circleutil.Circle
}

type rectangle struct {
	x  int
	y  int
	wh float64
	hh float64
}

// interface
type FiguresArea interface {
	area() float64
}

func (rec rectangle) area() float64 {
	return rec.wh * rec.hh
}

func (cir Circle) area() float64 {
	return math.Pi * math.Pow(cir.GetRadius(), 2)
}

func getArea(fig FiguresArea) float64 {
	return fig.area()
}

func main() {
	circle01 := circleutil.NewCir(1, 2, 3.6)
	fmt.Println(circle01.Area())

	circle02 := new(Circle)
	circle02.ChangeRadius(3.6)
	fmt.Println(circle02.area())

	rectangle01 := rectangle{1, 1, 2.5, 10}
	fmt.Println(rectangle01.area())

}
