package piscine

func checkPrintble(r rune) bool {
	if r >= 0 && r <= 31 {
		return true
	}
	return false
}

func IsPrintable(s string) bool {
	for _, v := range s {
		if checkPrintble(v) {
			return false
		}
	}
	return true
}
