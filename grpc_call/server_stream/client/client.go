package main

import (
	"context"
	"fmt"
	"github.com/jonnyGit81/coba/grpc_call/server_stream/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client...")
	fmt.Println()

	opts := grpc.WithInsecure()
	con, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("Error connecting: %v \n", err)
	}

	defer con.Close()
	c := documents.NewDocumentsClient(con)
	fetchDocuments(c)
}

// fetchDocuments function
func fetchDocuments(c documents.DocumentsClient) {
	// Initialize request message
	req := &documents.EmptyReq{}

	// Get the stream and err
	stream, err := c.GetDocuments(context.Background(), req)
	if err != nil {
		log.Fatalf("Error on GetDocuments: %v", err)
	}

	for {
		// Start receiving streaming messages
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error when receiving server response stream: %v", err)
		}
		log.Printf("Response from GetDocuments: %v", res.GetDocument())
	}
}
