package piscine

func checkNum(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	return false
}

func IsNumeric(s string) bool {
	for _, v := range s {
		if !checkNum(v) {
			return false
		}
	}
	return true
}
