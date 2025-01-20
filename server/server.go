package server

import (
	"bufio"
	"cli-database/cmd"
	"cli-database/database"
	"cli-database/lexer"
	"fmt"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	db         *database.Database
}

func NewServer(listenAddr string, db *database.Database) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		db:         db,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)

	if err != nil {
		return err
	}

	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("new connection to the server:", conn.RemoteAddr())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		if len(message) == 0 {
			continue
		}

		input := lexer.Tokenize(message)
		msg, err := cmd.Execute(input, s.db)
		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
		}
		conn.Write([]byte(msg))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("scanner error:", err)
	}
}
