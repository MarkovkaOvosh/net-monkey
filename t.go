package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// )

// var (
// 	clients  = make(map[string]net.Conn)
// 	leaving  = make(chan message)
// 	messages = make(chan message)
// )

// type message struct {
// 	text    string
// 	address string
// }

// func handle(conn net.Conn) {
// 	clients[conn.RemoteAddr().String()] = conn

// 	messages <- newMessage("joined.", conn)
// 	input := bufio.NewScanner(conn)
// 	for input.Scan() {
// 		messages <- newMessage(": "+input.Text(), conn)
// 	}
// 	delete(clients, conn.RemoteAddr().String())
// 	leaving <- newMessage(" has left. ", conn)
// 	conn.Close()
// }

// func newMessage(msg string, conn net.Conn) message {
// 	addr := conn.RemoteAddr().String()

// 	return message{
// 		text:    addr + msg,
// 		address: addr,
// 	}
// }

// func main() {
// 	listen, err := net.Listen("tcp", "localhost:4000")
// 	if err != nil {
// 		return
// 	}
// 	go broadcaster()
// 	for {
// 		conn, err := listen.Accept()
// 		if err != nil {
// 			log.Print(err)
// 			continue
// 		}
// 		go handle(conn)
// 	}
// }

// func broadcaster() {
// 	for {
// 		select {
// 		case msg := <-messages:
// 			for _, conn := range clients {
// 				if msg.address == conn.RemoteAddr().String() {
// 					continue
// 				}
// 				fmt.Fprintln(conn, msg.text+"dew")

// 			}
// 		case msg := <-leaving:
// 			for _, conn := range clients {
// 				fmt.Fprintln(conn, msg.text)
// 			}
// 		}
// 	}
// }
//
