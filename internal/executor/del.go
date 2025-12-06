package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
	"strconv"
)

type DelCommand struct{}

func (DelCommand) Execute(cli client.Client, args []string, store *store.Store) {
	if len(args) < 2 {
		SendMessage(cli.Conn, replies.InvalidArgs.Error())
		return
	}
	SendMessage(cli.Conn, strconv.Itoa(store.Del(args)))
}
