package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if (j == 0) && (i == 0 || i == y-1) {
					z01.PrintRune(65)
				} else if (i == 0 || i == y-1) && (j == x-1) {
					z01.PrintRune(67)
				} else if (i == 0 || i == y-1) || (j == 0 || j == x-1) {
					z01.PrintRune(66)
				} else {
					z01.PrintRune(32)
				}
			}
			z01.PrintRune('\n')
		}
	}
}
