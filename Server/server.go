package main

import (
	"context"
	Inc "github.com/SadlifeCaw/MiniProject3/Inc"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type Server struct {
	Inc.UnimplementedIncrementerServer
}

//slice of all known connections
var model = Inc.IncrementerModel{
	Counter:    0,
	User: "",
}

func main() {
	// Create listener tcp on port 3000
	//if port 3000 is occupied, use port 3001
	list, err := net.Listen("tcp", ":3000")
	ownPort := ":3000"

	if err != nil {

		list, err = net.Listen("tcp", ":3001")
		if err != nil {
			log.Fatalf("Failed to listen on port: %v", err)
		}
		ownPort = ":3001"
	}

	grpcServer := grpc.NewServer()
	Inc.RegisterIncrementerServer(grpcServer, &Server{})

	log.Println("Server is set up on port", ownPort)

	go IncrementerCounter()

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}

	//grpc listen and serve
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}

func IncrementerCounter() {
	//wait untill a increment is made
	for {
		if model.Counter > 0 {
			break
		}
	}
}

//grpc methods

func (s *Server) Increment(ctx context.Context, in *Inc.Request) (*Inc.Reply, error) {

	reply := Inc.Reply{}

	username := in.Username
	var replyMessage string

	// a better implementation would make sure each client has a unique ID rather than only rely on username
	var currentcount = model.Counter
	model.Counter = model.Counter + 1
	model.User = username

	replyMessage = "Counter is at " + strconv.Itoa(currentcount) + "."

	reply.ReplyMessage = replyMessage

	return &reply, nil
}

