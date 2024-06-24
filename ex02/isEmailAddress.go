package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Invalid argument\n")
		return 
	}
	re := regexp.MustCompile(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`)
	for i := 0; i < len(os.Args); i++ {
		if re.MatchString(os.Args[i]) {
			fmt.Printf("%s is a valid email address.\n", os.Args[i])
		} else {
			fmt.Printf("%s is NOT a valid email address.\n", os.Args[i])
		}
	}
}
