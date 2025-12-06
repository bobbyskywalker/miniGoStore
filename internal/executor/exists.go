package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
	"strconv"
)

type ExistsCommand struct{}

func (ExistsCommand) Execute(cli client.Client, args []string, store *store.Store) {
	if len(args) != 2 {
		SendMessage(cli.Conn, replies.InvalidArgs.Error())
		return
	}
	exists := store.Exists(args[1])
	SendMessage(cli.Conn, strconv.Itoa(exists))
}
