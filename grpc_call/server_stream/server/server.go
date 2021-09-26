package main

import (
	"github.com/jonnyGit81/coba/grpc_call/server_stream/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	documents.UnimplementedDocumentsServer
}

func main() {
	log.Println("Starting server...")

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()
	documents.RegisterDocumentsServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// container struct
type container struct {
	documents []*documents.Document
}

// GetDocuments function
func (*server) GetDocuments(req *documents.EmptyReq, stream documents.Documents_GetDocumentsServer) error {
	log.Println("GetDocuemnts function")

	// Initialize the container struct and call the initDocuments function
	// to get dummy data to send on the stream response message.
	docs := container{}.initDocuments()

	// Iterate over the documents
	for _, v := range docs {
		// Run some validation on each object
		if v.Size > 250 {
			// Create the response object
			res := &documents.GetDocumentsRes{
				Document: v,
			}

			// Use the stream object to send the response stream message
			stream.Send(res)

			// Sleep for a little bit..
			time.Sleep(1000 * time.Millisecond)
		}
	}
	return nil
}

// initDocuments function
func (c container) initDocuments() []*documents.Document {
	c.documents = append(c.documents, c.getDocument("Doc One", "nat", 345))
	c.documents = append(c.documents, c.getDocument("Doc Tow", "zip", 245))
	c.documents = append(c.documents, c.getDocument("Doc Three", "nat", 445))
	c.documents = append(c.documents, c.getDocument("Doc Four", "pid", 545))
	c.documents = append(c.documents, c.getDocument("Doc Five", "nat", 145))
	return c.documents
}

// getDocument function
func (c container) getDocument(name, documentType string, size int64) *documents.Document {
	return &documents.Document{
		Name:         name,
		DocumentType: documentType,
		Size:         size,
	}
}
