package piscine

func ThirdTimeIsACharm(str string) string {
	third := ""
	if len(str) < 3 || str == "" {
		return "\n"
	}
	for i := 2; i < len(str); i += 3 {
		third += string(str[i])

	}
	return third + "\n"
}
