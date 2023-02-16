package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type Users struct {
	Data    map[string]net.Conn
	Welcome string
	mu      sync.Mutex
}

type History struct {
	Cont string
	mu   sync.Mutex
}

type Message struct {
	User string
	Msg  string
	time time.Time
	conn net.Conn
}

var (
	mess   chan Message = make(chan Message)
	status chan Message = make(chan Message)
)

const (
	askName  = "[ENTER YOUR NAME]:"
	joinChat = " has joined our chat..."
	leftChat = " has left our chat..."
)

func main() {
	listen, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		fmt.Println(err)
		return
	}

	ping, err := os.ReadFile("cat.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	users := Users{
		Data:    map[string]net.Conn{},
		Welcome: string(ping),
		mu:      sync.Mutex{},
	}
	history := History{
		mu: sync.Mutex{},
	}

	go BroadCaster(&users, &history)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Fprintln(conn, err)
			conn.Close()
			continue
		}

		go Client(conn, &users, &history)
	}
}
