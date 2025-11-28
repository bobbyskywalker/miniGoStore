package client

import (
	"net"

	"github.com/segmentio/ksuid"
)

type Client struct {
	Conn net.Conn
	Id   string
}

func GenerateClientId() string {
	return (ksuid.New().String())
}
