package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Clinet(conn net.Conn, fileHistory, fileLog os.File) {
	defer fileHistory.Close()
	defer fileLog.Close()

	conn.Write([]byte("Welcome to the TCP chat!\n"))
	dog, err := os.ReadFile("dog.txt")
	if err != nil {
		loger(fmt.Sprintf("[COULDN'T OPEN FILE <logo.txt>][ADDRESS:%v]"), fileLog)
	}
	conn.Write(dog)
	var userName string

	newMsg := message{}

	for {
		conn.Write([]byte("[ENTER YOUR NAME]:"))
		inputName, err := bufio.NewReader(conn).ReadString('\n')
		userName, err = NameCheck(inputName)
		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
		}
		mu.Lock()
		clients[userName] = conn
		mu.Unlock()
		loger(fmt.Sprintf("[CLIENT JOIN IN THE CHAT][USER NAME:%v][ADDRESS:%v]", userName, conn.RemoteAddr().String()), fileLog)
		break
	}
	mu.Lock()
	history, err := os.ReadFile(fileHistory.Name())
	mu.Unlock()
	if err != nil {
		loger(fmt.Sprintf("[COULDN'T READ HISTORY FILE][ADDRESS:%v][ERROR:%s]", conn.RemoteAddr().String(), err.Error()), fileLog)
	}
	conn.Write(history)

	messages <- *newMsg.Add("\r"+userName+" has joined our chat..."+strings.Repeat(" ", 20), conn)
	conn.Write([]byte(template(userName)))

	input := bufio.NewScanner(conn)
	for input.Scan() {
		conn.Write([]byte(template(userName)))

		if !newMsg.Add(strings.TrimSpace(input.Text()), conn).Check() {
			netDate := *newMsg.Add(template(userName)+strings.TrimSpace(input.Text()), conn)
			mu.Lock()
			fileHistory.Write([]byte(netDate.text + "\n"))
			mu.Unlock()

			messages <- netDate
		}
	}

	loger(fmt.Sprintf("[CLIENT LEFT THE SERVER][USER NAME:%v][ADDRESS:%v]", userName, conn.RemoteAddr().String()), fileLog)
	mu.Lock()
	delete(clients, userName)
	mu.Unlock()
	status <- *newMsg.Add("\n"+userName+" has left our chat...", conn)
	conn.Close()
}
