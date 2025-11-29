package executor

import (
	"miniGoStore/internal/client"
)

type TtlCommand struct{}

func (TtlCommand) Execute(cli client.Client, args []string) {
}
