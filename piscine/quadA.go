package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if (j == 0 || j == x-1) && (i == 0 || i == y-1) {
					z01.PrintRune(111)
				} else if i == 0 || i == y-1 {
					z01.PrintRune(45)
				} else if j == 0 || j == x-1 {
					z01.PrintRune(124)
				} else {
					z01.PrintRune(32)
				}
			}
			z01.PrintRune('\n')
		}
	}
}
