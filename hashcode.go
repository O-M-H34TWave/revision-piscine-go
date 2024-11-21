package piscine

func HashCode(dec string) string {
	hash := ""
	for i := 0; i < len(dec); i++ {
		ascii := (len(dec) + int(dec[i])) % 127
		if ascii >= 32 && ascii <= 126 {
			hash += string(ascii)
		} else {
			ascii = (ascii % 95) + 32
			hash += string(ascii)
		}
	}
	return hash
}
