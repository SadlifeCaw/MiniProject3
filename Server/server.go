package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

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

func main() {

	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9081: %v", err)
	}

	grpcServer := grpc.NewServer()
	auction.RegisterAuctionServer(grpcServer, &Server{})

	log.Println("Server is set up on port 9081")

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}

	//grpc listen and serve
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}

	fmt.Println(model)
}

func (s *Server) Bid(ctx context.Context, in *auction.BidRequest) (*auction.BidReply, error) {
	username := in.Port
	bidAmountInt, err := strconv.Atoi(in.Message)
	var replyMessage string

	if err != nil {
		log.Fatalf("Error converting bid amount")
	}

	if bidAmountInt > int(model.HighestBid) {
		model.HighestBid = int(bidAmountInt)
		model.HighestBidder = username

		replyMessage = "You now have the highest bid by " + in.Message
	} else {
		highestBidString := strconv.Itoa(model.HighestBid)
		replyMessage = "Your bid did not go through. The highest current bid is " + highestBidString + " by " + model.HighestBidder
	}

	reply := auction.BidReply{
		ReplyMessage: replyMessage,
	}

	return &reply, nil
}
