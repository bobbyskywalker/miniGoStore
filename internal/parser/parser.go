package parser

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/executor"
	"strings"
)

func ParseCommand(cli client.Client, cmd []byte) {
	cmdStr := strings.TrimSpace(string(cmd))

	switch cmdStr {
	case "PING":
		executor.SendMessage(cli.Conn, "PONG")
	default:
	}
}
