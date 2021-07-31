package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("a byte is a number from 0 to 255")
	b1 := byte(0)
	b2 := byte(255)
	fmt.Println(b1, b2)
	fmt.Println("")
	fmt.Println("If you go below 0 or above 255 watch what happens")
	b1 = b1 - 1
	b2 = b2 + 1
	fmt.Println(b1, b2)
	fmt.Println("")
	fmt.Println("This is easier to understand if we look at the binary zeros and ones")
	fmt.Printf("%b %b\n", b1, b2)
	fmt.Println("")
	fmt.Println("A byte has 8 possible zeros or ones and 11111111 is the max")
	fmt.Println("If you add 1 it goes back to 00000000 or 0")
	fmt.Println("And if you are at 0 and subtrack 1, it goes to 11111111")
	fmt.Println("")
}
