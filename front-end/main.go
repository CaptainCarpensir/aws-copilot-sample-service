package main

import (
	"log"
	"net"
	"sync"
	"time"
)

const (
	UDPPort = 8081
	TCPPort = 8080
)

func handleUDPTraffic(wg *sync.WaitGroup) {
	defer wg.Done()

	udpServer, err := net.ListenUDP("udp", &net.UDPAddr{Port: UDPPort})
	if err != nil {
		log.Fatal(err)
	}

	defer udpServer.Close()
	log.Printf("Listening for UDP on port %d...\n", UDPPort)

	readBuffer := make([]byte, 1024)
	for {
		bytesRead, err := udpServer.Read(readBuffer)
		if err != nil {
			continue
		}
		msg := string(readBuffer[:bytesRead])
		log.Printf("Received UDP message: %q", msg)
	}
}

func handleTCPRequest(conn *net.TCPConn) {
	readBuffer := make([]byte, 1024)
	for {
		conn.SetReadDeadline(time.Now().Add(time.Minute))
		bytesRead, err := conn.Read(readBuffer)
		if err != nil {
			break
		}
		msg := string(readBuffer[:bytesRead])
		log.Printf("Received TCP message: %q", msg)
		_, err = conn.Write([]byte("pls"))
		if err != nil {
			break
		}
	}
	conn.Close()
	log.Println("Closed TCP connection")
}

func handleTCPTraffic(wg *sync.WaitGroup) {
	defer wg.Done()

	tcpServer, err := net.ListenTCP("tcp", &net.TCPAddr{Port: TCPPort})
	if err != nil {
		log.Fatal(err)
	}

	defer tcpServer.Close()
	log.Printf("Listening for TCP on port %d...\n", TCPPort)

	for {
		tcpConn, err := tcpServer.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Established TCP connection")

		go handleTCPRequest(tcpConn)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go handleTCPTraffic(&wg)
	go handleUDPTraffic(&wg)
	wg.Wait()
	log.Println("Finished listening")
}
