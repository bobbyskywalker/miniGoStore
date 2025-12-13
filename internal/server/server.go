package server

import (
	"errors"
	"io"
	"log/slog"
	"miniGoStore/internal/client"
	"miniGoStore/internal/parser"
	"miniGoStore/internal/store"
	"net"
	"sync/atomic"
)

const MsgBufSize = 1024

type Server struct {
	numClients int32
	storage    *store.Store
}

func NewServer() *Server {
	return &Server{
		numClients: 0,
		storage:    store.NewStore(),
	}
}

func (s *Server) StartServ(port string) {
	listener, err := net.Listen("tcp", ("localhost:" + port))
	if err != nil {
		slog.Error("Error:", slog.String("errormsg", err.Error()))
		return
	}
	defer listener.Close()

	s.storage.StartCleaner()

	slog.Info("miniGoStore Server is listening", slog.String("port", port))

	for {
		conn, err := listener.Accept()
		atomic.AddInt32(&s.numClients, 1)
		if err != nil {
			slog.Error("Error:", slog.String("errormsg", err.Error()))
			continue
		}
		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()

	cli := client.NewClient(conn)
	slog.Info("New client connected", slog.String("clientId", cli.Id))
	buf := make([]byte, MsgBufSize)

	for {
		nbytes, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF || errors.Is(err, net.ErrClosed) {
				slog.Info("Client disconnected", slog.String("clientId", cli.Id))
				atomic.AddInt32(&s.numClients, -1)
				return
			}
			slog.Error("Error:", slog.String("errormsg", err.Error()))
			return
		}
		slog.Debug("Received data", slog.String("clientId", cli.Id), slog.String("payload", string(buf[:nbytes])))
		parser.ParseCommand(cli, buf[:nbytes], s.storage)
	}
}
