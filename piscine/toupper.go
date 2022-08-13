package piscine

func ToUpper(s string) string {
	result := ""
	for _, v := range s {
		if v >= 97 && v <= 122 {
			result += string(v - 32)
		} else {
			result += string(v)
		}
	}
	return result
}
