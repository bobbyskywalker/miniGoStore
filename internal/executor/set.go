package executor

import (
	"miniGoStore/internal/client"
)

type SetCommand struct{}

func (SetCommand) Execute(cli client.Client, args []string) {

}
