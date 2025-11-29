package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/store"
)

type SetCommand struct{}

func (SetCommand) Execute(cli client.Client, args []string, store *store.Store) {

}
