package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func newMessage(msg string, conn net.Conn) message {
	addr := conn.RemoteAddr().String()
	return message{
		text:    addr + msg,
		address: addr,
	}
}

func loger(s string, file os.File) {
	str := fmt.Sprintf("[%s]%s\n", time.Now().Format("01-02-2006 15:04:05"), s)
	mu.Lock()
	defer mu.Unlock()
	_, err := file.WriteString(str)
	if err != nil {
		fmt.Print(fmt.Sprintf("[%s][COULDN'T WRITE TO FILE][ERROR:%s\n]", time.Now().Format("01-02-2006 15:04:05"), err))
	}
	fmt.Print(str)
}

func template(n string) string {
	return fmt.Sprintf("\r[%s][%s]:", time.Now().Format("01-02-2006 15:04:05"), n)
}

func (t message) Add(text string, c net.Conn) *message {
	return &message{
		text:    text,
		address: c.RemoteAddr().String(),
	}
}

func (t message) Check() bool {
	if t.text == "" {
		return true
	}
	for _, val := range t.text {
		if val < ' ' || val > '~' {
			return true
		}
	}
	return false
}

func NameCheck(str string) (string, error) {
	userName := strings.TrimSuffix(strings.TrimSpace(str), "\n")
	if len(userName) < 3 || len(userName) > 12 {
		return "", errors.New("Bad name! Len name must be more 2 and less 12 symbols\n")
	}
	for _, v := range userName {
		if v >= 'A' && v <= 'Z' || v >= 'a' && v <= 'z' {
			continue
		}
		return "", errors.New("Bad name! Name must has only latin alphabet\n")
	}
	mu.Lock()
	defer mu.Unlock()
	if _, ok := clients[userName]; ok {
		return "", errors.New("Name already has! Please try again.\n")
	}
	if len(clients) > 9 {
		return "", errors.New("Sever already full! Please try to connect later.\n")
	}
	return userName, nil
}
