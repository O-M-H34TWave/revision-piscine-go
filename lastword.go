package piscine

func LastWord(s string) string {
	last := ""
	// nbdaw mn last (len(s)-1 n3tiwha l i, then ndir i-- bach
	// nrj3o mn lor)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && len(last) > 0 {
			break
		}
		if s[i] != ' ' {
			// ndiro haka bach nzido lwords l beginning dial
			// last (ila drna last+=string(s[i]), adirna l
			// last word bl mgloub)
			last = string(s[i]) + last
		}

	}
	return last + "\n"
}
