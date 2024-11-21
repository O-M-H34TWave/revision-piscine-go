package piscine

func LastWord(s string) string {
	last := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			return ""
		}
	}
	return last
}
