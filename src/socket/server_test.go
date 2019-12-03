package sockettest

import (
	"log"
	"net"
	"testing"
)


func main ()  {
	socket, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Panicln(err)
	}

	defer socket.Close()

	for {
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

		conn.Write(buf[:size])
	}
}

func Test(t *testing.T){
	main()
}