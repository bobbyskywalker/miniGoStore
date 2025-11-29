package server

import (
	"io"
	"log"
	"miniGoStore/internal/client"
	"miniGoStore/internal/parser"
	"miniGoStore/internal/store"
	"net"
)

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

	log.Println("miniGoStore Server is listening on port " + port)

	for {
		conn, err := listener.Accept()
		s.numClients++
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()

	cli := client.Client{Conn: conn, Id: client.GenerateClientId()}
	log.Println("Client " + cli.Id + " connected")
	buf := make([]byte, 1024)

	for {
		nbytes, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("Client %s disconnected\n", cli.Id)
				s.numClients--
				return
			}
			log.Println("Error:", err)
			return
		}
		parser.ParseCommand(cli, buf[:nbytes], s.storage)
		log.Printf(cli.Id+": %s\n", buf[:nbytes])
	}
}
