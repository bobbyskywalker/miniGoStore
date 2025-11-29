package executor

import (
	"miniGoStore/internal/client"
)

type GetexCommand struct{}

func (GetexCommand) Execute(cli client.Client, args []string) {
}
