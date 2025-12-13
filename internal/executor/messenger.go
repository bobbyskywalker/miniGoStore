package executor

import (
	"log/slog"
	"net"
)

func SendMessage(conn net.Conn, msg string) {
	_, err := conn.Write([]byte(msg + "\n"))
	if err != nil {
		slog.Error("Error: cannot send message", slog.String("target", (conn.LocalAddr().String())))
	}
}
