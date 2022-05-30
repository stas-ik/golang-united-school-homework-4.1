package reverse_string

func ReverseString(input string) (output string) {
	reverseStringRunes := []rune(string)
	for i, j := 0, len(reverseStringRunes)-1; i < j; i, j = i+1, j-1 {
		reverseStringRunes[i], reverseStringRunes[j] = reverseStringRunes[j], reverseStringRunes[i]
	}
	return string(reverseStringRunes)
}
