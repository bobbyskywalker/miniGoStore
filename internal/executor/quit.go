package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/store"
)

type QuitCommand struct{}

func (QuitCommand) Execute(cli client.Client, args []string, store *store.Store) {
	cli.Conn.Close()
}
