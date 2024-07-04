package main

import (
	"fmt"
	"piscine"
	"strings"
)

func main() {
	fmt.Println(piscine.Compare("Hello!", "Hello!"))
	fmt.Println(strings.Compare("Hello!", "Hello!"))
	fmt.Println(piscine.Compare("Salut!", "lut!"))
	fmt.Println(strings.Compare("Salut!", "lut!"))
	fmt.Println(piscine.Compare("Ola!", "Ol"))
	fmt.Println(strings.Compare("Ola!", "Ol"))
	fmt.Println(piscine.Compare("", "Ol"))
	fmt.Println(strings.Compare("", "Ol"))
	fmt.Println(piscine.Compare("aaa", ""))
	fmt.Println(strings.Compare("aaa", ""))
	fmt.Println(piscine.Compare("123", "123456789"))
	fmt.Println(strings.Compare("123", "123456789"))
}
