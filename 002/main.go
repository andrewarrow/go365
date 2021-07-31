package main

import "fmt"

// go365 lesson 002: an uint16 is a number from 0 to 65535

func main() {
	fmt.Println("")
	fmt.Println("an uint16 is a number from 0 to 65535")
	i1 := uint16(0)
	i2 := uint16(65535)
	fmt.Println(i1, i2)
	fmt.Println("")
	fmt.Println("If you go below 0 or above 65535 watch what happens")
	i1 = i1 - 1
	i2 = i2 + 1
	fmt.Println(i1, i2)
	fmt.Println("")
	fmt.Println("This is easier to understand if we look at the binary zeros and ones")
	fmt.Printf("%b %b\n", i1, i2)
	fmt.Println("")
	fmt.Println("We had eight positions from lesson 001 with a byte and now we have double that")
	fmt.Println("")
	fmt.Println("32768 16384 8192 4096 2048 1024 512 256 128 64 32 16 8 4 2 1")
	fmt.Println("    1     1    1    1    1    1   1   1   1  1  1  1 1 1 1 1")
	fmt.Println("")
	fmt.Println("32768+16384+8192+4096+2048+1024+512+256+128+64+32+16+8+4+2+1=65535")
	fmt.Println("")
	fmt.Println("")
}
