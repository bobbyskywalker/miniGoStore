package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/errors"
	"miniGoStore/internal/store"
	"strconv"
	"time"
)

type SetCommand struct{}

func parsePassedTtl(i int, args []string, cli client.Client, result **time.Time, unit time.Duration) bool {
	if i+1 >= len(args) {
		SendMessage(cli.Conn, errors.InvalidArgs.Error())
		return false
	}
	nUnits, err := strconv.Atoi(args[i+1])
	if err != nil {
		SendMessage(cli.Conn, errors.InvalidArgs.Error())
		return false
	}
	t := time.Now().Add(time.Duration(nUnits) * unit)
	*result = &t
	return true
}

func (SetCommand) Execute(cli client.Client, args []string, store *store.Store) {
	var key string
	var val []byte
	var setOnExistent bool = false
	var setOnNonExistent bool = false
	var ttl *time.Time = nil
	var retrievePrevious = false

	if len(args) < 3 {
		SendMessage(cli.Conn, errors.InvalidArgs.Error())
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
