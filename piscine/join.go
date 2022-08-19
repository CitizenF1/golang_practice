package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i, v := range strs {
		if i != len(strs)-1 {
			result += v + sep
		} else {
			result += v
			break
		}
	}
	return result
}
