package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := []string(os.Args[1:])

	if isCorrect(args) {
		// table := fillTable(args)
	}

}

func isValid(table [9][9]rune, x, y int, currNum rune) bool {
	for i := 0; i < 9; i++ {
		if currNum == table[i][x] {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if currNum == table[y][j] {
			return false
		}
	}
	return true
}

func isSolve(table [9][9]rune) bool {
	if !isEmpty(table) {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if table[y][x] == '.' {
				for r := '1'; r <= '9'; r++ {
					if isValid(table, x, y, r) {
					}
				}
			}
		}

	}
}

func isEmpty(table [9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == '.' {
				return true
			}
		}
	}
	return false
}

func fillTable(args []string) [9][9]rune {
	var table [9][9]rune
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args); j++ {
			table[i][j] = rune(args[i][j])
		}
	}
	return table
}

func isCorrect(args []string) bool {
	// check if  out of the range
	if len(args) != 9 {
		return false
	}
	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for _, value := range args[i] {
			// check number range 1 - 9
			if value == 47 || value == 48 {
				return false
			} else if value < 46 || value > 57 {
				return false
			}
		}
	}
	return true
}

func printError() {
	err := "Error\n"
	for _, v := range err {
		z01.PrintRune(v)
	}
}
