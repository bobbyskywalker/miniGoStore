package server

import (
	"errors"
	"io"
	"log"
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
		log.Println("Error:", err)
		return
	}
	defer listener.Close()

	s.storage.StartCleaner()

	log.Println("miniGoStore Server is listening on port " + port)

	for {
		conn, err := listener.Accept()
		atomic.AddInt32(&s.numClients, 1)
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()

	cli := client.NewClient(conn)
	log.Println("Client " + cli.Id + " connected")
	buf := make([]byte, MsgBufSize)

	for {
		nbytes, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF || errors.Is(err, net.ErrClosed) {
				log.Printf("Client %s disconnected\n", cli.Id)
				atomic.AddInt32(&s.numClients, -1)
				return
			}
			log.Println("Error:", err)
			return
		}
		parser.ParseCommand(cli, buf[:nbytes], s.storage)
		log.Printf(cli.Id+": %s\n", buf[:nbytes])
	}
}
