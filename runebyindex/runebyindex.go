package runebyindex

import "log"

func RuneByIndex(s *string, i *int) (rune, error) 	 {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic recovered: %v", r)
		}
	}()
	var str string = *s
	var id int = *i
	return rune(str[id]), nil

}
