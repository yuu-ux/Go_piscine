package main

import (
	"ft"
	"os"
)

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}

func CountArray(s []string) int {
	var cnt int
	for i, _ := range s {
		cnt=i
	}
	return cnt
}

func main() {
	length := CountArray(os.Args)
	for 0 < length {
		PrintStr(os.Args[length])
		ft.PrintRune('\n')
		length--
	}
}
