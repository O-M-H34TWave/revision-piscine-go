package piscine

func PrintIf(str string) string {
	result := ""
	if len(str) >= 3 || len(str)==0 {
		result = "G\n"
	}else{
		return "Invalid Input\n"
	}
	return result
}
