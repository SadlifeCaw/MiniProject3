package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	t "time"

	auction "github.com/SadlifeCaw/MiniProject3/Auction"
	"google.golang.org/grpc"
)

var SliceOfServerPorts []string

var username string = ""

func main() {
	//hard coded to expect 3 servers on these ports
	SliceOfServerPorts = append(SliceOfServerPorts, ":3000", ":3001")

	//Read user input in terminal
	go ReadFromTerminal()

	fmt.Println("---------")
	fmt.Println("Bid on the auction by writing: 'Bid x', where 'x' is your bid amount")
	fmt.Println("Get the status on the auction by writing: 'Status'")
	fmt.Println("")

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

		//rn username is checked every time user types message, can be optimized
		if username == "" {
			username = clientMessage
			continue
		}

		clientMessage = strings.ToLower(clientMessage)
		splitString := strings.Fields(clientMessage)

		if splitString[0] == "bid" {

			bidAmount := splitString[1]
			BroadcastToServer(true, bidAmount)

		} else if splitString[0] == "status" {
			BroadcastToServer(false, "")
		} else {
			fmt.Println("Unknown command")
		}
	}
}

//use bool isBid, since the program only has two requests: Bid and Status
//not open for extension, but works
func BroadcastToServer(isBid bool, bidAmount string) {

	var SliceOfServerResponses []string

	//send request to all servers
	for _, port := range SliceOfServerPorts {
		var conn *grpc.ClientConn

		//the server has 1 second to respond, otherwise assumed to be dead
		//this will be very slow in case of a lot of dead servers, since the loop is not async
		conn, error := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))

		//if the server is dead, skip it
		if error != nil {
			continue
			//log.Fatalf("Could not connect: %s", error)
		}

		defer conn.Close()
		client := auction.NewAuctionClient(conn)

		//create context
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if isBid {
			//send bid request
			request := auction.BidRequest{
				Bid:      bidAmount,
				Username: username,
			}

			BidReply, err := client.Bid(ctx, &request)
			if err != nil {
				fmt.Println("Error while waiting for server reply (Bid)")
			}

			//save reply to slice
			SliceOfServerResponses = append(SliceOfServerResponses, BidReply.ReplyMessage)
		} else {
			//send status request
			request := auction.StatusRequest{}

			StatusReply, err := client.Status(ctx, &request)
			if err != nil {
				fmt.Println("Error while waiting for server reply (Status)")
			}

			SliceOfServerResponses = append(SliceOfServerResponses, StatusReply.ReplyMessage)
		}
	}

	//see if data is valid
	responsesAreValid := CheckReponseValidity(SliceOfServerResponses)

	if responsesAreValid {
		//if the data is valid, we can print any arbitrary value. pick 0 for bound safety
		fmt.Println(SliceOfServerResponses[0])
	} else {
		//if data is invalid, try to query servers again
		fmt.Println("Recieved different data from different servers. Retrying query")
		BroadcastToServer(isBid, bidAmount)
	}

}

//checks if all the server responses are equal
func CheckReponseValidity(ServerResponses []string) (isValid bool) {
	var length int = len(ServerResponses)
	var nextInt int = 0

	for i, message := range "ServerResponses" {
		nextInt = i + 1
		if nextInt >= length {
			break
		}

		if ServerResponses[i] != ServerResponses[nextInt] {
			return false
		}

		//why? because go is shit
		if message == 'p' {
			log.Fatalf("???")
		}
	}

	return true
}
