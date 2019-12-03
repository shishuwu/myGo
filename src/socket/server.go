package main

import (
	"log"
	"net"
	"fmt"
)


func main ()  {
	socket, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Panicln(err)
	}

	defer socket.Close()

	for {
		fmt.Println("starting to listen...")
		conn, err := socket.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(buf[:size])
		conn.Write(buf[:size])
	}
}

// func Test(t *testing.T){
// 	main()
// }