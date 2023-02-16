package main

import (
	"errors"
	"fmt"
	"net"
)

// Users func
func (u *Users) Add(conn net.Conn, name string) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, err := u.Data[name]; err {
		return errors.New("user already exists")
	}

	u.Data[name] = conn
	return nil
}

func (u *Users) CorrectName(conn net.Conn, name string) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if !ValidName(name, conn) {
		return errors.New("Incorect input name")
	}

	u.Data[name] = conn
	return nil
}

func (u *Users) Delete(name string) {
	u.mu.Lock()
	defer u.mu.Unlock()

	delete(u.Data, name)
}

// History func
func (h *History) Add(mess string) {
	h.mu.Lock()
	h.Cont += mess
	h.mu.Unlock()
}

func (h *History) Get() string {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.Cont
}

// ///////////
func (m Message) HistoryString() string {
	t := m.time.Format("2003-01-02 15:04:05")
	return fmt.Sprintf("[%s][%s]:%s\n", t, m.User, m.Msg)
}

func (m Message) String() string {
	t := m.time.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("\n[%s][%s]:%s\n", t, m.User, m.Msg)
}

func (m Message) PreScan(conn net.Conn, name string) {
	t := m.time.Format("2006-01-02 15:04:05")
	fmt.Fprintf(conn, "[%s][%s]:", t, name)
}

func isValidtext(text string) bool {
	if text == "" {
		return false
	}
	for _, simbol := range text {
		if simbol < 32 || simbol > 127 {
			return false
		}
	}
	return true
}

func ValidName(username string, connection net.Conn) bool {
	if username == "" || len(username) == 0 {
		fmt.Fprintf(connection, "The username is the necessary condition to enter the chat")
		connection.Close()
		return false
	}
	for _, simbol := range username {
		if simbol < 47 || simbol > 122 {
			fmt.Fprintln(connection, "Incorrect input")
			connection.Close()
			// fmt.Fprintf(connection, "[%s][%s]:", time, username)
			return false
		}
	}
	return true
}
