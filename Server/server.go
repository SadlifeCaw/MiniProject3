package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	auction "github.com/SadlifeCaw/MiniProject3/Auction"

	"google.golang.org/grpc"
)

type Server struct {
	auction.UnimplementedAuctionServer
}

//slice of all known connections
var sliceOfPorts []string
var model = auction.AuctionModel{
	HighestBid:    0,
	HighestBidder: "",
}

var ownPort string
var auctionOver bool
var auctionTime int = 60

func main() {

	fmt.Println("Choose port, type number in console:")
	fmt.Println("1: Port :3000")
	fmt.Println("2: Port :3001")

	var portChosen bool = false

	//let user choose server port
	for !portChosen {
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalf("Failed to read from console")
		}

		clientMessage = strings.Trim(clientMessage, "\r\n")

		if clientMessage == "1" {
			ownPort = ":3000"
			portChosen = true
		} else if clientMessage == "2" {
			ownPort = ":3001"
			portChosen = true
		} else {
			fmt.Println("1 or 2, not that hard")
		}
	}

	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ownPort)
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	grpcServer := grpc.NewServer()
	auction.RegisterAuctionServer(grpcServer, &Server{})

	log.Println("Server is set up on port", ownPort)

	go AuctionCounter()

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}

	//grpc listen and serve
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}

func AuctionCounter() {
	//wait untill a bid is made
	for {
		if model.HighestBid > 0 {
			break
		}
	}

	//decrease untill 0
	for auctionTime > 0 {
		time.Sleep(time.Second)
		auctionTime--
	}
	if auctionTime == 0 {
		auctionOver = true
	}
}

//grpc methods

func (s *Server) Bid(ctx context.Context, in *auction.BidRequest) (*auction.BidReply, error) {

	reply := auction.BidReply{}

	if auctionTime <= 0 {
		reply.ReplyMessage = "The auction has ended. You can no longer bid"

	} else {
		username := in.Username
		bidAmountInt, err := strconv.Atoi(in.Bid)
		var replyMessage string

		if err != nil {
			log.Fatalf("Error converting bid amount")
		}

		// a better implementation would make sure each client has a unique ID rather than only rely on username
		if bidAmountInt > int(model.HighestBid) {
			model.HighestBid = int(bidAmountInt)
			model.HighestBidder = username

			replyMessage = "You now have the highest bid by " + in.Bid + ". Time remaining " + strconv.Itoa(auctionTime) + " seconds"
		} else {
			highestBidString := strconv.Itoa(model.HighestBid)
			//it would make more sense to split this information in the replyMessage, and let the client handle formatting
			replyMessage = "Your bid did not go through. The highest current bid is " + highestBidString + " by " + model.HighestBidder + ". Time remaining " + strconv.Itoa(auctionTime) + " seconds"
		}

		reply.ReplyMessage = replyMessage
	}

	return &reply, nil
}

func (s *Server) Status(ctx context.Context, in *auction.StatusRequest) (*auction.StatusReply, error) {
	var stringToReturn string

	bidAmountInt := strconv.Itoa(model.HighestBid)

	if auctionOver {
		stringToReturn = "The auction has ended. The winner was " + model.HighestBidder + " who won with a bid of " + bidAmountInt
	}
	if auctionTime < 60 && !auctionOver {
		//auction started
		stringToReturn = "The auction is ongoing. The highest bid is " + bidAmountInt + " by " + model.HighestBidder + ". Time remaining " + strconv.Itoa(auctionTime) + " seconds"
	}
	if auctionTime == 60 && !auctionOver {
		stringToReturn = "The auction hasn't begun yet. Make a bid to start!"
	}

	reply := auction.StatusReply{
		ReplyMessage: stringToReturn,
	}

	return &reply, nil
}
