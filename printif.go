package piscine

func PrintIf(str string) string {
	result := ""
	if len(str) >= 3 || str == "" {
		result = "G\n"
	} else {
		result = "Invalid Input\n"
	}
	return result
}
