package main

import "os"

func main() {
	if len(os.Args) != 4 {
		return
	}
	input := os.Args[1]
	oldC := os.Args[2]
	newC := os.Args[3]
	res := ""
	for i := 0; i < len(input); i++ {
		if string(input[i]) == oldC {
			res += newC
		} else {
			res += string(input[i])
		}
	}
	println(res)
}
