package executor

import (
	"miniGoStore/internal/client"
)

type GetCommand struct{}

func (GetCommand) Execute(cli client.Client, args []string) {
}
