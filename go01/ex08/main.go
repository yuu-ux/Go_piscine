package main

import (
	"fmt"
	"piscine"
	"strconv"
)

func main() {
fmt.Println(piscine.BasicAtoi("12345"))
fmt.Println(strconv.Atoi("12345"))
fmt.Println(piscine.BasicAtoi("0000000012345"))
fmt.Println(strconv.Atoi("0000000012345"))
fmt.Println(piscine.BasicAtoi("000000"))
fmt.Println(strconv.Atoi("000000"))
}