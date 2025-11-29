package server

import (
	"log"
	"miniGoStore/internal/client"
	"miniGoStore/internal/parser"
	"miniGoStore/internal/store"
	"net"
)

type Server struct {
	numClients int32
	storage    store.Store
}

func StartServ(port string) {
	listener, err := net.Listen("tcp", ("localhost:" + port))
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer listener.Close()

	log.Println("miniGoStore Server is listening on port " + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error:", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	cli := client.Client{Conn: conn, Id: client.GenerateClientId()}
	log.Println("Client " + cli.Id + " connected")
	buf := make([]byte, 1024)

	for {
		nbytes, err := conn.Read(buf)
		if err != nil {
			log.Println("Error:", err)
			return
		}
		parser.ParseCommand(cli, buf[:nbytes])
		log.Printf(cli.Id+": %s\n", buf[:nbytes])
	}
}
