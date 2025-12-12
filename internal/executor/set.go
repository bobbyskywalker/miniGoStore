package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
	"time"
)

type SetCommand struct{}

func (SetCommand) Execute(cli client.Client, args []string, store *store.Store) {
	var key string
	var val []byte
	var setOnExistent bool = false
	var setOnNonExistent bool = false
	var ttl *time.Time = nil
	var retrievePrevious = false

	if len(args) < 3 {
		SendMessage(cli.Conn, replies.InvalidArgs.Error())
		return
	}

	key = args[1]
	val = []byte(args[2])

	for i := 3; i < len(args); i++ {
		switch args[i] {
		case "NX":
			setOnNonExistent = true
		case "XX":
			setOnExistent = true
		case "EX":
			ok := parsePassedTtl(i, args, cli, &ttl, time.Second)
			if !ok {
				return
			}
			i++
		case "PX":
			ok := parsePassedTtl(i, args, cli, &ttl, time.Millisecond)
			if !ok {
				return
			}
			i++
		case "GET":
			retrievePrevious = true
		}
	}
	v := store.Set(key, val, setOnExistent, setOnNonExistent, ttl, retrievePrevious)
	SendMessage(cli.Conn, string(v))
}
