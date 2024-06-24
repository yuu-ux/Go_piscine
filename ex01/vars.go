package main

import "fmt"

func main() {
fmt.Printf("%v\n", "42 is a string.")
fmt.Printf("%v\n", "42 is a uint.")
fmt.Printf("%v\n", "42 is an int.")
fmt.Printf("%v\n", "42 is a uint8.")
fmt.Printf("%v\n", "42 is an int16.")
fmt.Printf("%v\n", "42 is a uint32.")
fmt.Printf("%v\n", "42 is an int64.")
fmt.Printf("%v\n", "false is a bool.")
fmt.Printf("%v\n", "42 is a float32.")
fmt.Printf("%v\n", "42 is a float64.")
fmt.Printf("%v\n", "(42+0i) is a complex64.")
fmt.Printf("%v\n", "(42+21i) is a complex128.")
fmt.Printf("%v\n", "{} is a main.FortyTwo.")
fmt.Printf("%v\n", "[42] is a [1]int.")
fmt.Printf("%v\n", "map[42:42] is a map[string]int.")
fmt.Printf("%v\n", "0x0 is an *int.")
fmt.Printf("%v\n", "[] is a []int.")
fmt.Printf("%v\n", "<nil> is a chan int.")
fmt.Printf("%v\n", "<nil> is a <nil>.")
fmt.Printf("\n")
}
