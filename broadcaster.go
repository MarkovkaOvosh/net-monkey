package main

func BroadCaster() {
	for {
		select {
		case msg := <-messages:
			mu.Lock()
			// A b C
			for name, c := range clients {
				// name = B c = terminal- B           msg.Adr = A
				if msg.address == c.RemoteAddr().String() {
					c.Write([]byte(template(name)))
					// [02-17-171717 10:48:02][A]:
					continue
				}
				// [02-17-171717 10:52:46][A]: hello
				c.Write([]byte(msg.text + "\n"))
				// [02-17-171717 10:48:02][B]:
				c.Write([]byte(template(name)))
			}
			mu.Unlock()
		case msg := <-status:
			mu.Lock()
			for name, c := range clients {
				c.Write([]byte(msg.text + "\n"))
				c.Write([]byte(template(name)))
			}
			mu.Unlock()
		}
	}
}
