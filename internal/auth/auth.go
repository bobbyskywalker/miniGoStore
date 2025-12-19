package auth

import (
	"log/slog"
	"miniGoStore/internal/client"
	"miniGoStore/internal/executor"
	"miniGoStore/internal/replies"

	"golang.org/x/crypto/bcrypt"
)

func clearPassword(pass []byte) {
	for i := range pass {
		pass[i] = 0
	}
}

func HashPass(pass []byte) ([]byte, error) {
	defer clearPassword(pass)

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func comparePassAndHash(hash []byte, pass []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, pass) == nil
}

func HandleAuth(client *client.Client, hash []byte, pass []byte) {
	defer clearPassword(pass)

	if client.IsAuthenticated {
		executor.SendMessage(client.Conn, replies.RedundantAuthReply)
		return
	}

	if comparePassAndHash(hash, pass) {
		client.IsAuthenticated = true
		executor.SendMessage(client.Conn, replies.SuccessReply)
		slog.Info("Client registered", slog.String("clientID", client.Id))
	} else {
		executor.SendMessage(client.Conn, replies.InvalidPassword.Message)
		slog.Info("Client failed authentication", slog.String("clientID", client.Id))
	}
}
