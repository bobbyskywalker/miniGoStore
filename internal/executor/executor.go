package executor

import (
	"log"
	"net"
)

func SendMessage(conn net.Conn, msg string) {
	_, err := conn.Write([]byte(msg))
	if err != nil {
		log.Println("Error: cannot send message to: ", conn.LocalAddr())
	}
}
