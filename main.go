package main

import (
	"fmt"
	"github.com/zh-artykbay/hm3/atoi"
	"github.com/zh-artykbay/hm3/fibonacci"
	"github.com/zh-artykbay/hm3/itoa"
	"github.com/zh-artykbay/hm3/reverse"
	"github.com/zh-artykbay/hm3/runebyindex"
)

func main() {
	fmt.Println(itoa.Itoa(-332524353466))
	fmt.Println(atoi.Atoi("-34643436458368658"))
	fmt.Println(reverse.Reverse("dsfldslаывдла"))
	s := "sdf1sg"
	i := 10
	var st *string
	var id *int
	st = &s
	id = &i
	fmt.Println(runebyindex.RuneByIndex(st, id))
	generator := fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(generator(), " ")
	}

}

