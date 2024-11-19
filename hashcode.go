package piscine

func HashCode(dec string) string {
	code := ""
	for i := 0; i < len(dec); i++ {
		ascii := (len(dec) + int(dec[i])) % 127
		if ascii <32 || ascii > 126 {
			ascii+= 33
		}
		code += string(ascii)
	}
	return code
}
