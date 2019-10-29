package main

import (
	"fmt"

	"github.com/maciek13U/Go_tests/Simple01/circleutil"
)

// Provides access to the fields but not to the methods.
// type CircleHere circleutil.Circle

// Provides access to fields and methods
// type CircleHere struct {
// 	circleutil.Circle
// }

// alternative:
// type CircleHere struct {
//     *circleutil.Circle
// }

func main() {
	var1 := "abc"
	poinVar1 := &var1 // memory address

	fmt.Printf("Type: %T\n", poinVar1)
	fmt.Println(*poinVar1)

	*poinVar1 = "abc123"
	fmt.Println(var1)

	circle01 := circleutil.NewCir(1, 2, 3.6)
	fmt.Println(circle01)
	circle01.ChangeRadius(2.2)
	fmt.Println(circle01)
}
