package executor

import (
	"log/slog"
	"miniGoStore/internal/client"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
	"strconv"
)

type DelCommand struct{}

func (DelCommand) Execute(cli client.Client, args []string, store *store.Store) {
	defer slog.Info("Command completed", slog.String("clientId", cli.Id), "command", "DEL")

	if len(args) < 2 {
		SendMessage(cli.Conn, replies.InvalidArgs.Error())
		return
	}
	SendMessage(cli.Conn, strconv.Itoa(store.Del(args)))
}
