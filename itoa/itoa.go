package itoa


const digits = "0123456789abcdefghijklmnopqrstuvwxyz"
func Itoa(i int) string {
    var s string
	var isNegative bool
	if i == 0 {
		return "0"
	}
	if i < 0 {
		i = -i
		isNegative = true
	}
	for rem := i; i > 0; {
		rem = i % 10
		s += digits[rem : rem+1]
		i = i / 10
	}
	if isNegative {
		s += "-"
	}

	return reverse(s)
}

func reverse(s string) string {
	var reverse string
	bytes := []rune(s)
	for i := len(bytes) - 1; i >= 0; i-- {
		reverse += string(bytes[i])
	}
	return reverse
}
