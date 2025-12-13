package executor

import (
	"log/slog"
	"miniGoStore/internal/client"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
)

type PingCommand struct{}

func (PingCommand) Execute(cli client.Client, args []string, store *store.Store) {
	defer slog.Info("Command completed", slog.String("clientId", cli.Id), "command", "PONG")

	SendMessage(cli.Conn, replies.PongReply)
}
