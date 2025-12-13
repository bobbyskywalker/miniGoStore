package executor

import (
	"log/slog"
	"miniGoStore/internal/client"
	"miniGoStore/internal/store"
)

type QuitCommand struct{}

func (QuitCommand) Execute(cli client.Client, args []string, store *store.Store) {
	defer slog.Info("Command completed", slog.String("clientId", cli.Id), "command", "QUIT")

	cli.Conn.Close()
}
