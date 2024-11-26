package piscine

import "strconv"

func FromTo(from int, to int) string {
	res := ""
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}
	if from < to {
		for i := from; i <= to; i++ {
			if i != from {
				res += ", "
			}
			if i < 10 {
				res += "0"
			}
			res += strconv.Itoa(i)
		}
	} else {
		for i := from; i >= to; i-- {
			if i != from {
				res += ", "
			}
			if i < 10 {
				res += "0"
			}
			res += strconv.Itoa(i)
		}
	}
	res += "\n"
	return res
}
