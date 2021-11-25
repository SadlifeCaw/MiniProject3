package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	auction "github.com/SadlifeCaw/MiniProject3/Auction"

	"google.golang.org/grpc"
)

type Server struct {
	auction.UnimplementedAuctionServer
}

//slice of all known connections
var model = auction.AuctionModel{
	HighestBid:    0,
	HighestBidder: "",
	AuctionTime:   60, //init auction time to 60 seconds
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
	for model.AuctionTime > 0 {
		time.Sleep(time.Second)
		model.AuctionTime--
	}
}

//grpc methods

func (s *Server) Bid(ctx context.Context, in *auction.BidRequest) (*auction.BidReply, error) {

	reply := auction.BidReply{}

	if model.AuctionTime <= 0 {
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

			replyMessage = "You now have the highest bid by " + in.Bid + ". Time remaining " + strconv.Itoa(model.AuctionTime) + " seconds"
		} else {
			highestBidString := strconv.Itoa(model.HighestBid)
			//it would make more sense to split this information in the replyMessage, and let the client handle formatting
			replyMessage = "Your bid did not go through. The highest current bid is " + highestBidString + " by " + model.HighestBidder + ". Time remaining " + strconv.Itoa(model.AuctionTime) + " seconds"
		}

		reply.ReplyMessage = replyMessage
	}

	return &reply, nil
}

func (s *Server) Status(ctx context.Context, in *auction.StatusRequest) (*auction.StatusReply, error) {
	var stringToReturn string

	bidAmountInt := strconv.Itoa(model.HighestBid)

	if model.AuctionTime <= 0 {
		stringToReturn = "The auction has ended. The winner was " + model.HighestBidder + " who won with a bid of " + bidAmountInt
	}
	if model.AuctionTime < 60 && model.AuctionTime > 0 {
		//auction started
		stringToReturn = "The auction is ongoing. The highest bid is " + bidAmountInt + " by " + model.HighestBidder + ". Time remaining " + strconv.Itoa(model.AuctionTime) + " seconds"
	}
	if model.AuctionTime == 60 {
		stringToReturn = "The auction hasn't begun yet. Make a bid to start!"
	}

	reply := auction.StatusReply{
		ReplyMessage: stringToReturn,
	}

	return &reply, nil
}
