package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var (
	clients  = make(map[string]net.Conn)
	status   = make(chan message)
	messages = make(chan message)
	mu       sync.Mutex
)

type message struct {
	text    string
	address string
}

func main() {
	listen, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer listen.Close()

	h1, err := os.Create("history.txt")
	h2, err := os.Create("log.txt")
	loger(fmt.Sprintf("[SERVER WAS STARTED][PORT%v]", 4000), *h2)
	defer func() {
		if err := h1.Close(); err != nil {
			loger("[COULDN'T CLOSE FILE history.txt][ADDRESS:main]", *h2)
		}
		if err := h2.Close(); err != nil {
			loger("[COULDN'T CLOSE FILE log.txt][ADDRESS:main]", *h2)
		}
	}()
	go BroadCaster()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		loger(fmt.Sprintf("[CONNECT THE SERVER][ADDRESS:%v]", conn.RemoteAddr().String()), *h2)
		go Clinet(conn, *h1, *h2)
	}
}
