package executor

import (
	"miniGoStore/internal/client"
	"miniGoStore/internal/replies"
	"strconv"
	"time"
)

func parsePassedTtl(i int, args []string, cli client.Client, result **time.Time, unit time.Duration) bool {
	if i+1 >= len(args) {
		SendMessage(cli.Conn, replies.InvalidArgs.Error())
		return false
	}
	nUnits, err := strconv.Atoi(args[i+1])
	if err != nil {
		SendMessage(cli.Conn, replies.InvalidArgs.Error())
		return false
	}
	t := time.Now().Add(time.Duration(nUnits) * unit)
	*result = &t
	return true
}
