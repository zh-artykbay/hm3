package fibonacci


func Fibonacci() func() int {
	fst, snd := 0, 1
	return func() int {
		r := fst
		fst, snd = snd, fst+snd
		return r
	}
}
