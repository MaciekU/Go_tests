package main

import (
	"fmt"
	"strings"

	"github.com/maciek13U/Go_tests/Simple02/strutil"
)

// LetterInStringCount takes string and letter to count how many times
// it repeated (case sensitive!)
func LetterInStringCount(testedString string, letter string) int {
	lettersInString := strings.Count(testedString, letter)
	return lettersInString
}

// SumList adds elements from list of floats
func SumList(listSum []float64) float64 {
	var sum float64 = 0
	for _, elem := range listSum {
		sum += elem
	}
	return sum
}

func main() {
	var testString = "Aaaabbce1232234"
	var letter = "a"

	lettersInString := LetterInStringCount(testString, letter)

	fmt.Println(testString)
	fmt.Println(letter, ":", lettersInString)

	fmt.Println(strutil.Reverse("abc123"))

	// var strArray [3]string
	// strArray[0] = "123"
	// strArray[2] = "abc"
	strArray := [3]string{"123", "", "abc"}
	fmt.Println(strArray)

	// map[type]type
	var fileExtensions = map[string]string{
		"Python":     ".py",
		"C":          ".c",
		"C++":        ".cpp",
		"Text":       ".txt",
		"Executable": ".exe",
		"Golang":     ".go",
		"Java":       ".java",
	}

	if key, search := fileExtensions["Golang"]; search {
		fmt.Println(key)
	}

	testList := []float64{0, 2, 4, 0.1}
	fmt.Println(testList, " add=", SumList(testList))

}
