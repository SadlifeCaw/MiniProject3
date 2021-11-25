package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"

	t "time"

	auction "github.com/SadlifeCaw/MiniProject3/Auction"
)

func main() {
	// Creat a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//  Create new Client from generated gRPC code from proto
	client := auction.NewServerNodeClient(conn)

	fmt.Println("Type in your username")

	if err != nil {
		log.Fatalf("Failed to read from console")
	}

	//Read user input in terminal
	go ReadFromTerminal(ctx, client)

	//read from server
	go PrintBroadcastsFromServer(client)

	for {
		t.Sleep(1000 * t.Hour)
	}
}

func ReadFromTerminal(ctx context.Context, client auction.ServerNodeClient) {
	for {
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalf("Failed to read from console")
		}

		clientMessage = strings.Trim(clientMessage, "\r\n")

		if clientMessage == "bid" {
			bidRequest := auction.BidRequest{}

			client.Bid(ctx, &bidRequest)

		} else if "" == "status" {
			statusRequest := auction.StatusRequest{}
			client.Status(ctx, &statusRequest)
		}
	}
}

func PrintBroadcastsFromServer(client auction.ServerNodeClient) {

	for {
		messageToPrint := "Hello"
		fmt.Println(messageToPrint)
	}
}
