package main

import (
	"fmt"
	"strings"

	"github.com/maciek13U/Go_tests/Simple02/strutil"
)

// LetterInStringCount takes string and letter to count how many times it repeated
func LetterInStringCount(testedString string, letter string) int {
	lettersInString := strings.Count(testedString, letter)
	return lettersInString
}

func main() {
	var testString = "Aaaabbce1232234"
	var letter = "a"

	lettersInString := LetterInStringCount(testString, letter)

	fmt.Println(testString)
	fmt.Println(letter, ":", lettersInString)

	fmt.Println(strutil.Reverse("abc123"))
}
