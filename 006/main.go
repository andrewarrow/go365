package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// lesson 006: make your iot talkers act more real

// we want to test our server under realistic conditions;
// clients need to send more than just one byte every second.
// And we need more to connect than just one new connection
// every three seconds.

// we need:
// a random number of new connections per second,
// a random amount of data to send in each "payload",
// a random amount of time to continue sending before disconnecting,
// and a random amount of time to wait before re-connecting,

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
		delay := rand.Intn(1000)
		time.Sleep(time.Millisecond * delay) // this will create a random number per second
		// when delay is a small number like 10 or 20, or 40 over and over, and it randomly will be sometimes
		// you'll get a bunch of new connections back to back very fast
		// but when delay is larger, closer to 999 that will be almost a full one second delay
		// the end result is a nice realistic pattern of devices comming online with a brand new tcp connection
		// sometimes many within the same second
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
