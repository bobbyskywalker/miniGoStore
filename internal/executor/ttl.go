package executor

import (
	"log/slog"
	"miniGoStore/internal/client"
	"miniGoStore/internal/store"
	"strconv"
)

type TtlCommand struct{}

func (TtlCommand) Execute(cli client.Client, args []string, store *store.Store) {
	defer slog.Info("Command completed", slog.String("clientId", cli.Id), "command", "TTL")

	secondsLeft := store.CheckTtl(args[1])
	SendMessage(cli.Conn, (strconv.Itoa(secondsLeft)))
}
