package main

import "fmt"
import "net"

func main() {
	con, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		fmt.Println("error dialing network")
		return
	}
	defer con.Close()

	data := []byte("hello world")
	_, err = con.Write(data)
	if err != nil {
		fmt.Println("error writting data: ", err)
		return
	}

}
