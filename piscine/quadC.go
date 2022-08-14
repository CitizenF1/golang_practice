package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 {
					if j == 0 || j == x-1 {
						z01.PrintRune('A')
					} else {
						z01.PrintRune('B')
					}
				} else if i == y-1 {
					if j == 0 || j == x-1 {
						z01.PrintRune('C')
					} else {
						z01.PrintRune('B')
					}
				} else {
					if j == 0 || j == x-1 {
						z01.PrintRune('B')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}
