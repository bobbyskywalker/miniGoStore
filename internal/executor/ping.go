package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
)

type PingCommand struct{}

func (PingCommand) Execute(cli client.Client, args []string, store *store.Store) {
	SendMessage(cli.Conn, replies.PongReply)
}
