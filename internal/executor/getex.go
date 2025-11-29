package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/store"
)

type GetexCommand struct{}

func (GetexCommand) Execute(cli client.Client, args []string, store *store.Store) {
}
