package main

import (
	"log"
	"net"
	"os"

	auction "github.com/SadlifeCaw/MiniProject3/Auction"

	"google.golang.org/grpc"
)

type Server struct {
	auction.UnimplementedServerNodeServer
}

//slice of all known connections
var sliceOfPorts []string

func main() {

	//log to file, taken from https://dev.to/gholami1313/saving-log-messages-to-a-custom-log-file-in-golang-ce5
	LOG_FILE := "log.txt"

	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("-----------------------------------")

	log.Println("Logging to custom file")

	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9081")
	if err != nil {
		log.Fatalf("Failed to listen on port 9081: %v", err)
	}

	grpcServer := grpc.NewServer()
	auction.RegisterServerNodeServer(grpcServer, &Server{})

	log.Println("Server is set up on port 9081")

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}

	//grpc listen and serve
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}
