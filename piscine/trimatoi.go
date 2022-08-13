package piscine

func TrimAtoi(s string) int {
	res := 0
	var minus, firstDigit bool
	for _, v := range s {
		if v == '-' && !firstDigit {
			minus = true
		}
		if v >= 48 && v <= 57 {
			firstDigit = true
			res = res*10 + int(v) - '0'
		}
	}
	if minus {
		return -res
	}
	return res
}
