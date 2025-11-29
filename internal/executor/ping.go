package executor

import (
	"miniGoStore/internal/client"
)

type PingCommand struct{}

func (PingCommand) Execute(cli client.Client, args []string) {
	SendMessage(cli.Conn, "PONG")
}
