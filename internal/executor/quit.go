package executor

import (
	"miniGoStore/internal/client"
)

type QuitCommand struct{}

func (QuitCommand) Execute(cli client.Client, args []string) {
}
