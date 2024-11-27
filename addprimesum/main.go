package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	pos := os.Args[1]
	result := 0
	if pos == "" {
		return
	}
	i, _ := strconv.Atoi(pos)
	// ndiro var jdid ndiro fih ybda b 1 w ybda ydir i%x ta ywsel l i
	// i in this case hya arg li ghandiroha f main
	// donc x%1 == 0 x%2!=0....x%i==0 then ay appendihom
	for x := 1; x <= i; x++ {
		if i%x == 0 {
			result += i
		}
	}
	fmt.Println(result)
}
