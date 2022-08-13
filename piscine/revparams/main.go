package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	r := os.Args
	for i := len(r) - 1; i > 0; i-- {
		for _, v := range r[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
