package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"
)

// lesson 007: quality vs quantity in fmt.Print

var statsMutex = sync.Mutex{} // try removing this and any refernce to it, run it and see error
var statsConnection = map[string]int64{}
var statsLength = map[string]int64{}
var statsData = map[string]int64{}

func main() {
	rand.Seed(time.Now().UnixNano())

	thing, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	go DisplayStats()
	go MakeConnections()
	for {
		conn, err := thing.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		remote := conn.RemoteAddr()
		ip := remote.String()
		statsMutex.Lock()
		statsConnection[ip] = time.Now().Unix()
		statsMutex.Unlock()
		go ReadInZerosAndOnes(ip, conn)
	}
}

func MakeConnections() {
	for {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			fmt.Println(err)
			return
		}
		go SendZerosAndOnes(conn)
		delay := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * delay)
	}
}

func SendZerosAndOnes(conn net.Conn) {
	for {
		size := rand.Intn(10) + 1
		payload := make([]byte, size)
		rand.Read(payload)

		conn.Write(payload)
		delay := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * delay)

		if rand.Intn(100) > 98 {
			conn.Close()
			break
		}
	}
}

func ReadInZerosAndOnes(ip string, conn net.Conn) {
	for {
		payload := make([]byte, 10)
		n, err := conn.Read(payload)
		if err != nil {
			if err == io.EOF {
				statsMutex.Lock()
				delta := time.Now().Unix() - statsConnection[ip]
				statsLength[ip] = delta
				statsMutex.Unlock()
				break
			}
		}

		statsMutex.Lock()
		statsData[ip] += int64(n)
		statsMutex.Unlock()
	}
}

func DisplayStats() {
	for {
		time.Sleep(time.Millisecond * 3000)
		statsMutex.Lock()
		fmt.Printf("Total      : %d\n", len(statsData))
		fmt.Printf("Total Died : %d\n", len(statsLength))
		fmt.Println("")
		statsMutex.Unlock()
	}
}
