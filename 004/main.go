package main

import (
	"fmt"
	"time"
)

// go365 lesson 004: maps and goroutines

func main() {

	go func() {
		fmt.Println("got here 1.")
	}()

	go func() {
		fmt.Println("got here 2.")
	}()

	fmt.Println("got here 3.")  // this is on the main goroutine
	time.Sleep(time.Second * 1) // try removing this line, notice the change?
}

// run this over and over. notice the order is not predictable.
// there are three goroutines here, the main one you start in, and
// two more we create.

// then rename main above to main2 and rename this main2 below to main
// this is a easy way to switch back and forth which main function you run

func main2() {
	m := map[string]int{}

	go func() {
		m["a-thing"] = 1
		m["any_string_really"] = 17
		m[" "] = 1000
		m["even a \" quote"] = -200
	}()

	go func() {
		for {
			for key, value := range m {
				fmt.Println(key, value)
				m[key]++
			}
			time.Sleep(time.Second * 1) // try removing this line, notice the change?
		}
	}()

	time.Sleep(time.Second * 5) // try removing this line, notice the change?
}
