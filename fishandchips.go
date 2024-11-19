package piscine

func FishAndChips(n int) string {
	result := ""
	if n%2 == 0 {
		result = "fish"
	}
	if n%3 == 0 {
		result = "chips"
	}
	if n%2 == 0 && n%3 == 0 {
		result = "fish and chips"
	}
	if n < 0 {
		result = "error: number is negative"
	} else if n%2 != 0 && n%3 != 0 {
		result = "error: non divisible"
	}
	return result
}
