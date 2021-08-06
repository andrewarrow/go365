package main

import (
	"fmt"
	"sync"
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
// this is ain easy way to switch back and forth which main function you run

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

// run this and notice what a map does, it maps a key and a value together.
// you can access the same map from two differrent goroutines but it
// isn't safe. You might not see the problem right away
// but given enough time, you will get the error `fatal error: concurrent map writes`

// a way to fix this is with a mutex lock

func main3() {
	mutex := sync.Mutex{}
	m := map[string]int{}

	go func() {
		mutex.Lock()
		m["abc"] = 1
		mutex.Unlock()
	}()

	go func() {
		for {
			mutex.Lock()
			m["abc"]++
			fmt.Println(m["abc"])
			mutex.Unlock()
			time.Sleep(time.Second * 1) // try removing this line, notice the change?
		}
	}()

	time.Sleep(time.Second * 5) // try removing this line, notice the change?
}

func main4() {
}
