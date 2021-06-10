package reverse

func Reverse(s string) string {
	var reverse string
	bytes := []rune(s)
	for i := len(bytes) - 1; i >= 0; i-- {
		reverse += string(bytes[i])
	}
	return reverse
}
