package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/store"
)

type Command interface {
	Execute(cli client.Client, args []string, store *store.Store)
}

var Commands = map[string]Command{
	"PING":   PingCommand{},
	"QUIT":   QuitCommand{},
	"SET":    SetCommand{},
	"GET":    GetCommand{},
	"GETEX":  GetexCommand{},
	"DEL":    DelCommand{},
	"EXISTS": ExistsCommand{},
	"TTL":    TtlCommand{},
}
