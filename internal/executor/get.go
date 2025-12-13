package executor

import (
	"log/slog"
	"miniGoStore/internal/client"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
)

type GetCommand struct{}

func (GetCommand) Execute(cli client.Client, args []string, store *store.Store) {
	defer slog.Info("Command completed", slog.String("clientId", cli.Id), "command", "GET")

	if len(args) != 2 {
		SendMessage(cli.Conn, replies.InvalidArgs.Error())
		return
	}
	key := args[1]
	value, ok := store.Get(key)
	if !ok {
		SendMessage(cli.Conn, replies.NotFoundReply)
		return
	}
	SendMessage(cli.Conn, string(value.Value))
}
