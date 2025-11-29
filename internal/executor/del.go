package executor

import (
	"miniGoStore/internal/client"
)

type DelCommand struct{}

func (DelCommand) Execute(cli client.Client, args []string) {
}
