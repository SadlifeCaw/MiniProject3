package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	t "time"

	auction "github.com/SadlifeCaw/MiniProject3/Auction"
	"google.golang.org/grpc"
)

var sliceOfServerPorts []string

var username string = ""

func main() {
	//hard coded to expect 3 servers on these ports
	sliceOfServerPorts = append(sliceOfServerPorts, ":3000")

	//Read user input in terminal
	go ReadFromTerminal()

	fmt.Println("Bid on the auction by writing: 'Bid x', where 'x' is your bid amount")
	fmt.Println("Get the status on the auction by writing: 'Status'")

	for {
		t.Sleep(1000 * t.Hour)
	}
}

func ReadFromTerminal() {
	fmt.Println("Please enter your username")

	for {
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalf("Failed to read from console")
		}

		clientMessage = strings.Trim(clientMessage, "\r\n")
		splitString := strings.Fields(clientMessage)

		//rn username is checked every time user types message, can be optimized
		if username == "" {
			username = clientMessage
		}

		if splitString[0] == "bid" {

			bidAmount := splitString[1]
			BroadcastToServer(true, bidAmount)

		} else if splitString[0] == "status" {
			BroadcastToServer(false, "")
		}
	}
}

func BroadcastToServer(isBid bool, bidAmount string) {
	for _, port := range sliceOfServerPorts {

		//fmt.Println(isBid, port)

		var conn *grpc.ClientConn
		conn, error := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
		if error != nil {
			log.Fatalf("Could not connect: %s", error)
		}

		// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
		defer conn.Close()

		//  Create new Client from generated gRPC code from proto
		client := auction.NewAuctionClient(conn)
		//create context
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if isBid {

			request := auction.BidRequest{
				Message: bidAmount,
				Port:    username,
			}

			BidReply, err := client.Bid(ctx, &request)
			if err != nil {
				fmt.Println("Error while waiting for server reply (Bid)")
			}

			fmt.Println(BidReply.ReplyMessage)
		}
	}
}

func PrintBroadcastsFromServer(client auction.AuctionClient) {

	for {
		messageToPrint := "Hello"
		fmt.Println(messageToPrint)
	}
}
