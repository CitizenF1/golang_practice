package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := []string(os.Args[1:])
	if len(args) == 1 {
		args = SplitWhiteSpaces(args[0])
	}
	if isCorrect(args) {
		table := fillTable(args)
		if backtrack(&table) {
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					if j != 9 {
						z01.PrintRune(table[i][j])
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}
			z01.PrintRune('\n')
		} else {
			printError()
		}
	} else {
		printError()
	}
}

func checkDuplicate(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func isValid(table [9][9]rune, column, row int, currNum rune) bool {
	// check coll
	for i := 0; i < 9; i++ {
		if currNum == table[i][column] {
			return false
		}
	}
	// check row
	for j := 0; j < 9; j++ {
		if currNum == table[row][j] {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		counter := [10]int{}
		for j := 0; j < 9; j++ {
			if int(table[i][j])-'0' >= 1 {
				counter[int(table[i][j]-'0')]++
			}
		}
		if checkDuplicate(counter) {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		counter := [10]int{}
		for j := 0; j < 9; j++ {
			if int(table[j][i])-'0' >= 1 {
				counter[int(table[j][i]-'0')]++
			}
		}
		if checkDuplicate(counter) {
			return false
		}
	}
	// check square
	a, b := column/3, row/3
	for i := 3 * a; i < 3*(a+1); i++ {
		for j := 3 * b; j < 3*(b+1); j++ {
			if currNum == table[j][i] {
				return false
			}
		}
	}
	return true
}

func backtrack(table *[9][9]rune) bool {
	if !isEmpty(table) {
		return true
	}
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if table[row][column] == '.' {
				for r := '1'; r <= '9'; r++ {
					// if cell == '.' try fill number 1 - 9 and backtrack
					if isValid(*table, column, row, r) {
						table[row][column] = r
						if backtrack(table) {
							return true
						}
					}
					table[row][column] = '.'
				}
				return false
			}
		}
	}
	return false
}

func isEmpty(table *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == '.' {
				return true
			}
		}
	}
	return false
}

func isCorrect(args []string) bool {
	count := 0
	// check range
	if len(args) != 9 {
		return false
	}
	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if args[i][j] != '.' {
				count++
			}
			if (args[i][j] <= 48 || args[i][j] > 57) && !(args[i][j] == '.') {
				return false
			}
		}
	}
	// fmt.Println(count)
	return count >= 17
}

func printError() {
	err := "Error\n"
	for _, v := range err {
		z01.PrintRune(v)
	}
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

func SplitWhiteSpaces(s string) []string {
	var res []string
	word := ""
	for i, v := range s {
		if !isSeparate(v) {
			word += string(v)
		}
		if isSeparate(v) && word != "" || len(s) == i+1 && word != "" {
			res = append(res, word)
			word = ""
		}
	}
	return res
}

func isSeparate(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}
