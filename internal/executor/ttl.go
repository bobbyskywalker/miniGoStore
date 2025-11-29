package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/store"
)

type TtlCommand struct{}

func (TtlCommand) Execute(cli client.Client, args []string, store *store.Store) {
}
