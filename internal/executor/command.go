package executor

import "miniGoStore/internal/client"

type Command interface {
	Execute(cli client.Client, args []string)
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
