package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	asc := true
	desc := true
	for i := 1; i < len(a); i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			asc = false
		}
	}
	for i := 1; i < len(a); i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			desc = false
		}
	}
	return asc || desc
}

func isSort(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
