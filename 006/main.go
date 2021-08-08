package main

import (
	"fmt"
	"io"
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
// a random rate per second of sending these payloads,
// and a random amount of time to continue sending before disconnecting,

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
		delay := time.Duration(rand.Intn(1000))
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
		size := rand.Intn(10) + 1 // add 1 to make range 1 to 10 vs 0 to 9
		payload := make([]byte, size)
		rand.Read(payload)

		conn.Write(payload)
		delay := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * delay) // this will also create a random number per second

		if rand.Intn(100) > 95 { // about a 5% chance this connection dies
			conn.Close()
			break
		}
	}
}

func ReadInZerosAndOnes(ip string, conn net.Conn) {
	for {
		payload := make([]byte, 10)  // we don't know the size, but we know max will be 10
		n, err := conn.Read(payload) // the Read function is smart enough to only place the number of available bytes in to payload
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		fmt.Printf("Receieved %08b from client %s\n", payload[0:n], ip)
	}
}

// run and leave running for about two to three minutes.
// try and follow a certain client that stays connected for a long time
// think about questions like, what is the average length of time a client stays connected?
// what is the average amount of data a client sends over its life?
// what are some stats on which clients sent how much data

// try and make the output going to fmt.Print less in volume but better in quality
