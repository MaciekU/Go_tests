package strutil

func Reverse(str string) string {
	chars := []rune(str)

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 { // no i++, j-- !!!
		chars[i], chars[j] = chars[j], chars[i]
	}

	return str + " reversed: " + string(chars)
}
