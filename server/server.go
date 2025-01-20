package server

import (
	"bufio"
	"cli-database/cmd"
	"cli-database/database"
	"cli-database/lexer"
	"fmt"
	"net"
)

type Message struct {
	from    string
	payload string
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	Msgch      chan Message
	db         *database.Database
}

func NewServer(listenAddr string, db *database.Database) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		Msgch:      make(chan Message, 1000),
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
	go s.messageLoop()

	<-s.quitch
	close(s.Msgch)

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
		if len(message) > 0 {
			s.Msgch <- Message{
				from:    conn.RemoteAddr().String(),
				payload: message,
			}
		}
		conn.Write([]byte(""))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("scanner error:", err)
	}
}

func (s *Server) messageLoop() {
	for msg := range s.Msgch {

		input := lexer.Tokenize(msg.payload)
		msg, err := cmd.Execute(input, s.db)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(msg)
		}
	}
}
