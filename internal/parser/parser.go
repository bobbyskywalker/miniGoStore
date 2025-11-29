package parser

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/executor"
	"strings"
)

func ParseCommand(cli client.Client, cmd []byte) {
	cmdStr := strings.TrimSpace(string(cmd))

	tokens := strings.Split(cmdStr, " ")
	name := tokens[0]

	if handler, ok := executor.Commands[name]; ok {
		handler.Execute(cli, tokens)
		return
	}
	executor.SendMessage(cli.Conn, "ERR: unknown command")
}
