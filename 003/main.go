package main

import "fmt"

// go365 lesson 003: know your ints

func main() {
	fmt.Println("")
	fmt.Println("byte   : 0 255")
	fmt.Println("uint16 : 0 to 65535")
	fmt.Println("uint32 : 0 to 4294967295")
	fmt.Println("uint64 : 0 to 18446744073709551615")
	fmt.Println("int16  : -32768 to 32767")
	fmt.Println("int32  : -2147483648 to 2147483647")
	fmt.Println("int64  : -9223372036854775808 to 9223372036854775807")
	fmt.Println("")
	fmt.Println("When you just say int it will default to either")
	fmt.Println("")
	fmt.Println("  int32 or int64")
	fmt.Println("")
	fmt.Println("Depending on the machine you are on. This is known as Platform dependent")

	test := uint64(18446744073709551615) // try chaning this number
	fmt.Printf("%b\n", test)

}
