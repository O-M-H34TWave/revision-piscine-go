package piscine

func CheckNumber(arg string) bool {
	for _, e := range arg {
		if e >= '0' && e <= '9'{
			return true
		}
	}
	return false
}
