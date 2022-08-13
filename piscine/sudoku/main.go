package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := []string(os.Args[1:])

	if isCorrect(args) {
		table := fillTable(args)
		if isSolve(table) {
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					if j != 8 {
						z01.PrintRune(rune(table[i][j]))
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(rune(table[i][j]))
					}
				}
				z01.PrintRune(10)
			}
		} else {
			printError()
		}
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
	a, b := x/3, y/3

	for i := 3 * a; i < 3*(a+1); i++ {
		for j := 3 * b; j < 3*(b+1); j++ {
			if currNum == table[j][i] {
				return false
			}
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
						table[y][x] = r
						if isSolve(table) {
							return true
						}
					}
					table[y][x] = '.'
				}
				return false
			}
		}
	}
	return false
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
