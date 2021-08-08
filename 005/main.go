package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// lesson 005: building something real

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

// as you run this and see output like

/*

making a connection
connection established: [::1]:8080
[::1]:62435
Receieved [00001111] from client [::1]:62435
Receieved [10110000] from client [::1]:62435
Receieved [11110001] from client [::1]:62435
making a connection
connection established: [::1]:8080
[::1]:62436
Receieved [01101011] from client [::1]:62436
Receieved [00100000] from client [::1]:62435
Receieved [10100000] from client [::1]:62435
Receieved [00101111] from client [::1]:62436
Receieved [01000111] from client [::1]:62435
Receieved [11111010] from client [::1]:62436
making a connection
Receieved [11101011] from client [::1]:62436
Receieved [01110011] from client [::1]:62435
[::1]:62439
connection established: [::1]:8080
Receieved [01100011] from client [::1]:62439

*/

// Let's interpret what this is saying:

// we see "making a connection"
// then "connection established" on port 8080
// 62435 is the part of the ip address of the client connecting
// then we get three messages from this client, 00001111,10110000,11110001
// then "making a connection" and "connection established" again on port 8080
// ah it's somebody new 62436
// notice from this new client we get 01101011,00101111,11111010
// but also notice I had to find these mixed in with more data from 62435
// 00100000,10100000,01000111
// They are sending these data simultaneously
// Then we get 62439 and the output stops here

// add more fmt.Println statements in this program in various
// places and change the output, run it over and over until
// you feel like 006
