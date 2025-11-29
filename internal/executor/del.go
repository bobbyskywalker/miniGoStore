package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/store"
)

type DelCommand struct{}

func (DelCommand) Execute(cli client.Client, args []string, store *store.Store) {
}
