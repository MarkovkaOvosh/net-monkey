package main

import "fmt"

func BroadCaster(u *Users, h *History) {
	for {
		select {
		case msg := <-mess:
			h.Add(msg.HistoryString())
			u.mu.Lock()
			for name, conn := range u.Data {
				if name != msg.User {
					fmt.Fprint(conn, msg)
				}
				msg.PreScan(conn, name)

			}
			u.mu.Unlock()
		case stat := <-status:
			h.Add(stat.User + stat.Msg + "\n")
			u.mu.Lock()
			for name, conn := range u.Data {
				if name != stat.User {
					fmt.Fprint(conn, "\n", stat.User+stat.Msg, "\n")
					stat.PreScan(conn, name)
				}
			}
			u.mu.Unlock()
		}
	}
}
