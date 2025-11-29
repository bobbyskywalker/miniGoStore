package executor

import (
	"miniGoStore/internal/client"
)

type ExistsCommand struct{}

func (ExistsCommand) Execute(cli client.Client, args []string) {
}
