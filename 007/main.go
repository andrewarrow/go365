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

var statsMutex = sync.Mutex{}
var statsConnection = map[string]int64{}
var statsData = map[string]int64{}

func main() {
	rand.Seed(time.Now().UnixNano())

	thing, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	go MakeConnections()
	for {
		conn, _ := thing.Accept()
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
				statsMutex.Unlock()
				fmt.Printf("%s Lasted %d seconds.\n", ip, delta)
				break
			}
		}

		statsMutex.Lock()
		statsData[ip] += int64(n)
		statsMutex.Unlock()
	}
}
