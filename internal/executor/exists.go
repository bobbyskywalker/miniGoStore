package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/store"
)

type ExistsCommand struct{}

func (ExistsCommand) Execute(cli client.Client, args []string, store *store.Store) {
}
