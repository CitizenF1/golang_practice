package piscine

func Map(f func(int) bool, a []int) []bool {
	var arr []bool
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			arr = append(arr, true)
		} else {
			arr = append(arr, false)
		}
	}
	return arr
}
