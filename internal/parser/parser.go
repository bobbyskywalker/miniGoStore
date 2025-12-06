package parser

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/executor"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
	"strings"
)

func ParseCommand(cli client.Client, cmd []byte, store *store.Store) {
	cmdStr := strings.TrimSpace(string(cmd))

	tokens := strings.Split(cmdStr, " ")
	name := strings.ToUpper(tokens[0])

	if handler, ok := executor.Commands[name]; ok {
		handler.Execute(cli, tokens, store)
		return
	}
	executor.SendMessage(cli.Conn, replies.UnknownCommand.Error())
}
