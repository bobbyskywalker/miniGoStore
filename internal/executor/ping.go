package executor

import (
	"miniGoStore/internal/client"
	server_replies "miniGoStore/internal/server_replies/replies"
	"miniGoStore/internal/store"
)

type PingCommand struct{}

func (PingCommand) Execute(cli client.Client, args []string, store *store.Store) {
	SendMessage(cli.Conn, server_replies.PongReply)
}
