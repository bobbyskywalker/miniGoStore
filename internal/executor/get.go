package executor

import (
	"log"
	"miniGoStore/internal/client"
	"miniGoStore/internal/replies"
	"miniGoStore/internal/store"
)

type GetCommand struct{}

func (GetCommand) Execute(cli client.Client, args []string, store *store.Store) {
	if len(args) != 2 {
		SendMessage(cli.Conn, replies.InvalidArgs.Error())
		return
	}
	key := args[1]
	log.Println("Retrieving value for key: " + key + " for client: " + cli.Id)
	value, ok := store.Get(key)
	if !ok {
		log.Printf("Value for key: %s not found", key)
		SendMessage(cli.Conn, replies.NotFoundReply)
		return
	}
	SendMessage(cli.Conn, string(value.Value))
}
