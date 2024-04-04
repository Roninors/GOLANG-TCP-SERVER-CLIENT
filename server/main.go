package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

var (
	mut           sync.Mutex
	listOfClients []*net.TCPConn
)

func main() {
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		fmt.Println("error occured: ", err)
		return
	}

	defer listener.Close()
	fmt.Println("server is listening on port 4000")
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection: ")
			continue
		}
		tcpConn := con.(*net.TCPConn)
		tcpConn.SetKeepAlive(true)
		tcpConn.SetNoDelay(true)
		mut.Lock()
		listOfClients = append(listOfClients, tcpConn)
		mut.Unlock()
		go handleClient(con)
	}

}

func handleClient(con net.Conn) {
	defer con.Close()
	buffer := make([]byte, 512)
	for {
		message, err := con.Read(buffer)
		if err != nil {
			if err == io.EOF {

				fmt.Println("client disconnected ")
				return
			}
			fmt.Println("error reading buffer: ", err)
			return
		}

		fmt.Printf("Recieved, %s\n", strings.TrimSpace(string(buffer[:message])))
		broadcast(string(buffer[:message]))
	}
}

func broadcast(message string) {
	for _, client := range listOfClients {
		_, err := client.Write([]byte(message))
		if err != nil {
			log.Printf("Error broadcasting message: %+v", err.Error())
		}
	}
}
