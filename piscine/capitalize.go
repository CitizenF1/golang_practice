package piscine

func checkA(r rune) bool {
	if r >= 48 && r <= 57 || r >= 97 && r <= 122 || r >= 65 && r <= 90 {
		return true
	}
	return false
}

func Capitalize(s string) string {
	first := true
	runes := []rune(s)
	for i := range runes {
		if checkA(runes[i]) && first {
			first = false
			if runes[i] >= 97 && runes[i] <= 122 {
				runes[i] -= 32
			}
		} else if runes[i] >= 65 && runes[i] <= 90 {
			runes[i] += 32
		} else if !checkA(runes[i]) {
			first = true
		}
	}
	return string(runes)
}
