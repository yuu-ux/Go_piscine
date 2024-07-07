package main

import (
	"ft"
	"os"
)

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func Sort(s1 *string, s2 *string) {
	*s1, *s2 = *s2, *s1
}

func CountArray(array []string) int {
	var cnt int
	for i, _ := range array {
		cnt = i
	}
	return cnt
}
func main() {
	length := CountArray(os.Args)

	for i := 0; i < length; i++ {
		for j := length; 0 < j; j-- {
			if os.Args[j-1] > os.Args[j] {
				Sort(&os.Args[j-1], &os.Args[j])
			}
		}
	}
	for _, s := range os.Args[1:] {
		PrintStr(s)
	}
}
