package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, v := range args {
		for _, r := range v {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func bubbleSort(arr []rune) []rune {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			var tmp rune
			if arr[i] > arr[j] {
				tmp = rune(arr[i])
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}
