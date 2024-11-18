package piscine

func RetainFirstHalf(str string) string {
	first := ""
	if len(str) == 1 {
		return str
	}
	for e := 0; e < len(str)/2; e++ {
		first += string(str[e])
	}
	return first
}
