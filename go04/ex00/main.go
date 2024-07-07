package main

import (
	"os"
	"ft"
)

func findLastsrash(s string) int {
	indexSrash := 0
	for i, r := range s {
		if r == '/' {
			indexSrash = i + 1
		}
	}
	return indexSrash
}

func main() {
	indexSrash := findLastsrash(os.Args[0])
	os.Stdout.Write([]byte(os.Args[0][indexSrash:]))
	ft.PrintRune('\n')
}
