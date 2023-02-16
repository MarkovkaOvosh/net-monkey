package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func Client(conn net.Conn, u *Users, h *History) {
	defer conn.Close()
	scan := bufio.NewScanner(conn)

	var username string

	conn.Write([]byte(u.Welcome))
	conn.Write([]byte(askName))

	if scan.Scan() {
		username = scan.Text()
	} else {
		return
	}
	if err := u.Add(conn, username); err != nil {
		conn.Write([]byte(err.Error()))
		return
	}
	if err := u.CorrectName(conn, username); err != nil {
		conn.Write([]byte(err.Error()))
		return
	}
	defer u.Delete(username)

	status <- Message{
		User: username,
		Msg:  joinChat,
		time: time.Now(),
	}

	conn.Write([]byte(h.Get()))

	msg := Message{
		User: username,
		time: time.Now(),
		conn: conn,
	}

	msg.PreScan(conn, username)
	for scan.Scan() {
		text := strings.Trim(scan.Text(), " ")
		if !isValidtext(text) {
			fmt.Fprintln(conn, "The empty messages are prohibited")
			fmt.Fprintf(conn, "[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), username)
			continue
		}
		msg.Msg = scan.Text()
		msg.time = time.Now()

		mess <- msg
	}

	status <- Message{
		User: username,
		Msg:  leftChat,
		time: time.Now(),
	}
}
