package main

import (
	"fmt"
	"log/slog"
	"miniGoStore/internal/auth"
	"miniGoStore/internal/logger"
	"miniGoStore/internal/server"
	"os"
)

func main() {
	args := os.Args[1:]
	port := ""

	switch len(args) {
	case 1:
		port = "8080"
	case 2:
		port = args[1]
	default:
		fmt.Println("Valid exec.: ./miniGoStore <password> [port]")
		os.Exit(1)
	}

	hash, err := auth.HashPass([]byte(args[0]))
	if err != nil {
		panic(err)
	}
	logger.InitLogger(slog.LevelInfo)
	s := server.NewServer(hash)
	s.StartServ(port)
}
