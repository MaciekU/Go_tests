package main

import (
	"fmt"

	"github.com/maciek13U/Go_tests/Simple01/circleutil"
)

func add(a, b int) int {
	return a + b
}

func main() {

	var str1 string
	// var str1 = "Zażółć gęślą jaźń!" automatic type!
	// str1 := "Zażółć gęślą jaźń!"
	str1 = "Zażółć gęślą jaźń!"

	calculateAdd := add(2, 3)

	fmt.Println(str1, "\nOperation add:", calculateAdd)

	r := circleutil.NewCir(2, 4, 2.5)
	fmt.Println("Circle x,y and radius: ", r)
	fmt.Println("Area of circle:", r.Area())
}
