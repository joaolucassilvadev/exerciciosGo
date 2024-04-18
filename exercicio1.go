package main

import (
	"fmt"
)

func main() {
	d := "joao"
	c := "lucas"

	nome(d, c)
}

func nome(d, c string) bool {
	i := 0
	l := 0
	for k := range d {
		fmt.Printf("%x\n", k)

		i++
	}

	for j := range c {
		fmt.Printf("%x\n", j)

		l++
	}

	if i > 1 && l > 1 {
		return true
	} else {
		return false
	}
}
