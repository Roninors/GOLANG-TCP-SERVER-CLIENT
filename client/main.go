package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	con, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		fmt.Println("error dialing network")
		return
	}
	defer con.Close()
	connReader := bufio.NewReader(con)
	go func() {
		for {
			serverReply, err := connReader.ReadString('\n')
			if err != nil {
				fmt.Println("error occured during reading server reply: ", err)
				return
			}
			fmt.Println("server reply: ", strings.TrimSpace(serverReply))
		}
	}()
	for {
		fmt.Println("Send server a message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error when reading message: ", err)
			return
		}
		data := []byte(message)
		_, err = con.Write(data)
		if err != nil {
			fmt.Println("error writting data: ", err)
			return
		}

	}

}
