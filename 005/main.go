package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// go365 lesson 005: building something real

// picture every IoT device in the entire world making a tcp connection
// to this program. Millions of them, connecting, sending data (zeros and ones)
// and then disconnecting at some point and repeating this over and over.

// in order to stress test and practice running our server to support this
// we need to write both sides of the program. The server and the client.

func main() {
	rand.Seed(time.Now().UnixNano()) // do this at start just once

	thing, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	go MakeConnections()
	for {
		conn, _ := thing.Accept()
		remote := conn.RemoteAddr()
		fmt.Println(remote)
		go ReadInZerosAndOnes(remote.String(), conn)
	}
}

func MakeConnections() {
	for {
		fmt.Println("making a connection")
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("connection established:", conn.RemoteAddr())
		go SendZerosAndOnes(conn)
		time.Sleep(time.Second * 3)
	}
}

func SendZerosAndOnes(conn net.Conn) {
	for {
		oneByte := make([]byte, 1)
		rand.Read(oneByte) // this will be 8 zeros and ones something like 01001100 or 10100111, etc.

		conn.Write(oneByte)
		time.Sleep(time.Second * 1)
	}
}

func ReadInZerosAndOnes(ip string, conn net.Conn) {
	for {
		oneByte := make([]byte, 1)
		conn.Read(oneByte)
		fmt.Printf("Receieved %08b from client %s\n", oneByte, ip)
		// 08 in between the % and b means
		// make it 8 in length and add 0 to the front is they are missing
		// try it with just %b and see the difference
	}
}
