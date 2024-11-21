package piscine

func CamelToSnakeCase(s string) string {
	result := ""
	if s == "" {
		return ""
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' && i != 0 {
			result += "_" + string(s[i])
		} 
	}
	return result
}
