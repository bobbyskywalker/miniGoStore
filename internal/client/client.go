package client

import (
	"net"

	"github.com/segmentio/ksuid"
)

type Client struct {
	Conn net.Conn
	Id   string
}

func generateClientId() string {
	return (ksuid.New().String())
}

func NewClient(conn net.Conn) Client {
	return Client{
		Conn: conn,
		Id:   generateClientId(),
	}
}
