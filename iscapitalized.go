package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	startOfWord := true

	for i := 0; i < len(s); i++ {
		if startOfWord {
			if s[i] >= 'a' && s[i] <= 'z' {
				return false
			}
			startOfWord = false
		}
		if s[i] == ' ' {
			startOfWord = true
		}
	}
	return true
}
