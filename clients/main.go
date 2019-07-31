package main

import (
	"context"

	"time"

	"log"

	c "app/contact"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:8080"
	defaultName = "world"
)

// func listName(c c.HelloServiceClient) {
// 	stream, err := c.ListName(context.Background(), &c.ListNameRequest{})
// 	if err != nil {
// 		log.Fatalf("%v", err)
// 	}

// 	for {
// 		msg, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("%v.ListFeatures(_) = _, %v", stream, err)
// 		}
// 		log.Println(msg.GetName())
// 	}
// }

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := c.NewContactServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// r, err := client.AddContact(ctx, &c.AddContactReq{Name: "e", Phone: "Phone"})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }

	// r, err := client.GetContact(ctx, &c.GetContactReq{Id: 1})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("%s", r)

	r, err := client.SearchContact(ctx, &c.SearchContactReq{Q: "a"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for _, contact := range r.Contact {
		log.Printf("%v", contact)
	}
}
