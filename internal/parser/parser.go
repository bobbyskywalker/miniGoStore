package parser

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/executor"
	"strings"
)

func ParseCommand(cli client.Client, cmd []byte) {
	cmdStr := strings.TrimSpace(string(cmd))

	/* TODO: tokenize and parse
	 * implement a strategy pattern for this */

	switch cmdStr {
	case "PING":
		executor.SendMessage(cli.Conn, "PONG")
	case "SET":
	case "GET":
	case "GETEX":
	case "DEL":
	case "EXISTS":
	case "TTL":
	case "QUIT":
	default:
	}
}
