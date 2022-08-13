package piscine

func check(r rune) bool {
	if r >= 65 && r <= 90 {
		return true
	}
	return false
}

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if !check(rune(s[i])) {
			return false
		}
	}
	return true
}
