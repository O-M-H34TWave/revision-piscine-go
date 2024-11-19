package piscine

func PrintIfNot(str string) string {
	result := ""
	if len(str) < 3{
		result = "G\n"
	}else {
		return "Invalid Input\n"
	}
	return result
}
