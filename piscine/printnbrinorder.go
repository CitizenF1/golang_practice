package piscine

import "github.com/01-edu/z01"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			var tmp int
			if arr[i] > arr[j] {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}

func PrintNbrInOrder(n int) {
	if n > 0 {
		var arr []int
		for n != 0 {
			arr = append(arr, n%10)
			n /= 10
		}
		newArr := bubbleSort(arr)
		for _, v := range newArr {
			z01.PrintRune('0' + rune(v))
		}
	} else {
		z01.PrintRune('0')
	}
}
