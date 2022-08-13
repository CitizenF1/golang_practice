package piscine

func checkLower(r rune) bool {
	if r >= 97 && r <= 122 {
		return true
	}
	return false
}

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if !checkLower(rune(s[i])) {
			return false
		}
	}
	return true
}
