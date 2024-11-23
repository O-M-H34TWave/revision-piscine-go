package piscine

func FirstWord(s string) string {
	res := ""
	if s == "" {
		return "\n"
	}
	for i := 0; i < len(s); i++ {
		// nchofo wach current char is a space w wach deja douzna 3la lword lwl (len(s)>0)
		if s[i] == ' ' && len(s) > 0 {
			break
		}
		// ila current char is not a space, n3mro res
		if s[i] != ' ' {
			res += string(s[i])
		}
	}
	return res + "\n"
}
