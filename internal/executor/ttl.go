package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/store"
	"strconv"
)

type TtlCommand struct{}

func (TtlCommand) Execute(cli client.Client, args []string, store *store.Store) {
	secondsLeft := store.CheckTtl(args[1])
	SendMessage(cli.Conn, (strconv.Itoa(secondsLeft)))
}
