package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
	"time"
)

type GetexCommand struct{}

func (GetexCommand) Execute(cli client.Client, args []string, store *store.Store) {
	var ttl *time.Time = nil
	var persist bool = false

	var key string = args[1]

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "EX":
			ok := parsePassedTtl(i, args, cli, &ttl, time.Second)
			if !ok {
				return
			}
		case "PX":
			ok := parsePassedTtl(i, args, cli, &ttl, time.Millisecond)
			if !ok {
				return
			}
			i++
		case "PERSIST":
			persist = true
		}
	}

	if ttl != nil && persist {
		SendMessage(cli.Conn, replies.SyntaxError.Message)
		return
	}

	v, ok := store.SetEx(key, ttl, persist)
	if !ok {
		SendMessage(cli.Conn, replies.NotFoundReply)
	}
	SendMessage(cli.Conn, string(v.Value))
}
