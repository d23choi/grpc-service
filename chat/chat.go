package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct{}

// Implement the rpc method SayHello() defined in chat.proto
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello from the server"}, nil
}