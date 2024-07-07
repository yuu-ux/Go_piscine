package main

import (
	"os"
	"ft"
)

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}
func main() {
	for _, s := range os.Args[1:] {
		PrintStr(s)
		ft.PrintRune('\n')
	}
}
