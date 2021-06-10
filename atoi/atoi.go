package atoi

import "errors"

func Atoi(s string) (int, error) {

	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if len(s) < 1 {
			return 0, errors.New("Invalid Syntax")
		}
	}

	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0, errors.New("Invalid Syntax")
		}
		n = n*10 + int(ch)
	}
	if s0[0] == '-' {
		n = -n
	}
	return n, nil
}
