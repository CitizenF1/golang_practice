package piscine

func ToLower(s string) string {
	result := ""
	for _, v := range s {
		if v >= 65 && v <= 90 {
			result += string(v + 32)
		} else {
			result += string(v)
		}
	}
	return result
}
