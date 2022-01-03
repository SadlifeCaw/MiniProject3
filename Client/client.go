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

	Inc "github.com/SadlifeCaw/MiniProject3/Inc"
	"google.golang.org/grpc"
)

var SliceOfServerPorts []string

var username string = ""

func main() {
	//hard coded to expect 2 servers on these ports
	SliceOfServerPorts = append(SliceOfServerPorts, ":3000", ":3001")

	//Read user input in terminal
	go ReadFromTerminal()
	fmt.Print("╒═════ INCREMENTER BASE ═════╕\n│      CTRL + C to leave     │\n│   Write Inc to Increment   │\n└────────────────────────────┘\n\n")

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

		if clientMessage == "inc" {
			BroadcastToServer()

		} else {
			fmt.Println("Unknown command")
		}
	}
}

//not open for extension, but works
func BroadcastToServer() {

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
		client := Inc.NewIncrementerClient(conn)

		//create context
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

			//send increment request
			request := Inc.Request{
				Username: username,
			}

			Reply, err := client.Increment(ctx, &request)
			if err != nil {
				fmt.Println("Error while waiting for server reply")
			}

			//save reply to slice
			SliceOfServerResponses = append(SliceOfServerResponses, Reply.ReplyMessage)
	}

	//see if data is valid
	responsesAreValid := CheckReponseValidity(SliceOfServerResponses)

	if responsesAreValid {
		//if the data is valid, we can print any arbitrary value. pick 0 for bound safety
		fmt.Println(SliceOfServerResponses[0])
	} else {
		//if data is invalid, try to query servers again
		fmt.Println("Recieved different data from different servers. Retrying query")
		BroadcastToServer()
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
