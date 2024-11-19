package piscine

func RetainFirstHalf(str string) string {
	first := ""
	for e := 0; e <= len(str)/2; e++ {
		if len(str) == 0 {
			return ""
		}
		if len(str) == 1 {
			return str
		} else {
			first += string(str[e])
		}
	}
	return first
}
