package main

import (
	"fmt"
	"io"
	"net"
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
		go handleClient(con)
	}

}

func handleClient(con net.Conn) {
	defer con.Close()
	buffer := make([]byte, 512)
	for {
		message, err := con.Read(buffer)
		if err != nil {
            if err == io.EOF{
                             
			fmt.Println("client disconnected ")
            return
            }
			fmt.Println("error reading buffer: ", err)
			return
		}

		fmt.Printf("Recieved, %s\n", buffer[:message])
	}
}
