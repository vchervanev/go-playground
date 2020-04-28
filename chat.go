package main

import (
	"bufio"
	"log"
	"net"
)

func handle(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var s server = server{connections: make(map[int]net.Conn)}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	handle(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		id := s.register(conn)
		go handleConnection(id, conn)
	}
}

func handleConnection(id int, c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		s.onMessage(c, input.Text())
	}
}

type server struct {
	count       int
	connections map[int]net.Conn
}

func (s *server) register(c net.Conn) (result int) {
	result = s.count
	s.connections[result] = c
	s.count++
	return
}
func (s *server) onMessage(c net.Conn, message string) {
	for _, con := range s.connections {
		if c != con {
			con.Write([]byte(message))
			con.Write([]byte("\n"))
		}
	}
}
