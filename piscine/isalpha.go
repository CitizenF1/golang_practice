package piscine

func checkAlpha(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	} else if r >= 97 && r <= 122 {
		return true
	} else if r >= 65 && r <= 90 {
		return true
	}
	return false
}

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !checkAlpha(rune(s[i])) {
			return false
		}
	}
	return true
}
