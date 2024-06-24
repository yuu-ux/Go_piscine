package main

import (
	"fmt"
	"os"
	"strconv"
)
// go run createStairs.go 4
func main() {
	j := 0
	i := 0
	arg_num, _ := strconv.Atoi(os.Args[1])
	if (arg_num >= 0 && arg_num <= 10000) {
		for i < arg_num {
			for 0 < j {
				fmt.Printf("*")
				j++
			}
			i++
			fmt.Printf("\n")
		}
	}
}
