package main

import (
	"fmt"
	"miniGoStore/internal/server"
	"os"
)

func main() {
	args := os.Args[1:]
	port := ""

	switch len(args) {
	case 0:
		port = "8080"
	case 1:
		port = args[0]
	default:
		fmt.Println("Valid exec.: ./miniGoStore [port]")
		os.Exit(1)
	}

	server.StartServ(port)
}
