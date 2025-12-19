package parser

import (
	"log/slog"
	"miniGoStore/internal/auth"
	"miniGoStore/internal/client"
	"miniGoStore/internal/executor"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
	"strings"
)

func ParseCommand(cli *client.Client, cmd []byte, store *store.Store, hash []byte) {
	slog.Debug("Start parse", slog.String("clientId", cli.Id))
	cmdStr := strings.TrimSpace(string(cmd))

	tokens := strings.Split(cmdStr, " ")
	name := strings.ToUpper(tokens[0])

	if name == "PASS" {
		auth.HandleAuth(cli, hash, []byte(tokens[1]))
		return
	}

	if cli.IsAuthenticated {
		if handler, ok := executor.Commands[name]; ok {
			handler.Execute(*cli, tokens, store)
			return
		}
		executor.SendMessage(cli.Conn, replies.UnknownCommand.Error())
		slog.Debug("End parse: unknown command", slog.String("clientId", cli.Id))
	} else {
		executor.SendMessage(cli.Conn, replies.UnauthorizedReply)
	}
}
