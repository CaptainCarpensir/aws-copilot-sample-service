# sasdfdasdf
package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	TCPAddr = "127.0.0.1:8080"
	UDPAddr = "127.0.0.1:8081"
)

func SendUDPTraffic(wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.Dial("udp", UDPAddr)
	if err != nil {
		log.Fatal("UDP Connection failed")
	}
	defer conn.Close()

	start := time.Now()
	for {
		time_passed := time.Now().Sub(start)
		msg := fmt.Sprintf("Time since script started: %s", time_passed.Round(time.Second))
		conn.Write([]byte(msg))
		log.Printf("Sent UDP message: %q\n", msg)
		time.Sleep(time.Second)
	}
}

func SendTCPTraffic(wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.Dial("tcp", TCPAddr)
	if err != nil {
		log.Fatal("TCP Connection failed")
	}
	defer conn.Close()
	log.Printf("Established TCP connection with %s\n", conn.RemoteAddr())

	start := time.Now()
	for {
		conn.SetReadDeadline(time.Now().Add(time.Minute))
		time_passed := time.Now().Sub(start)
		msg := fmt.Sprintf("Time since script started: %s", time_passed.Round(time.Second))
		conn.Write([]byte(msg))
		log.Printf("Sent TCP message: %q\n", msg)
		// readBuffer := make([]byte, 1024)
		// bytesRead, err := conn.Read(readBuffer)
		// if err != nil {
		// 	log.Fatal(err)
		// 	break
		// }
		// msg = string(readBuffer[:bytesRead])
		// log.Printf("Received TCP message: %q", msg)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go SendTCPTraffic(&wg)
	// go SendUDPTraffic(&wg)
	wg.Wait()
}
