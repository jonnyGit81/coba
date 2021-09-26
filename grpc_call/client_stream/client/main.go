package main

import (
	"context"
	"fmt"
	"github.com/jonnyGit81/coba/grpc_call/client_stream/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client Stream...")
	fmt.Println()

	opts := grpc.WithInsecure()
	con, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("Error connecting: %v \n", err)
	}

	defer con.Close()
	c := books.NewBooksClient(con)
	validateBooks(c)
}

// container struct
type container struct {
	books []*books.ValidationReq
}

// validateBooks function
func validateBooks(c books.BooksClient) {
	// Initialize the container struct and call the initBooks function
	// to get dummy data to send on the request message.
	req := container{}.initBooks()

	// Get the stream and err
	stream, err := c.ValidateBooks(context.Background())
	if err != nil {
		log.Fatalf("FATAL Error on ValidateBooks: %v", err)
	}

	// Iterate over the request message
	for _, v := range req {
		// Start making streaming requests by sending
		// each book object inside the request message
		fmt.Println("Client streaming request: \n", v)
		fmt.Println()
		stream.Send(v)
		time.Sleep(500 * time.Millisecond)
	}

	// Once the for loop finishes, the stream is closed
	// and get the response and a potential error
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("FATAL Error when closing the stream and receiving the response: %v", err)
	}

	// Print the response errors message
	fmt.Println()
	fmt.Printf("Validation errors: %v \n", res.Errors)
}

// initBooks function
func (c container) initBooks() []*books.ValidationReq {
	c.books = append(c.books, c.getBook("1", "Book 1", "This is a really good book about history", "John Phill", "2.5", 395, 2010))
	c.books = append(c.books, c.getBook("2", "Book 2", "Improve your communication skills", "Carl Matz", "3.5", 425, 2008))
	c.books = append(c.books, c.getBook("3", "Book 3", "Movies and TV shows", "Carl Matz", "2.9", 225, 2015))
	c.books = append(c.books, c.getBook("3", "Bo", "Cookies", "Carl Matz", "2.2", 275, 2011))
	c.books = append(c.books, c.getBook("4", "Book 5", "Learn more about animals and nature", "John Phill", "2.7", 455, 2018))
	c.books = append(c.books, c.getBook("4", "Book 6 with long title", "Machine learning", "John Phill", "1.0", 375, 2016))
	c.books = append(c.books, c.getBook("5", "Book", "10 good ideas for decorating your house", "Maty Metzer", "3.2", 250, 2007))
	return c.books
}

// getBook function
func (c container) getBook(id, title, desc, author, edition string, pages, year int64) *books.ValidationReq {
	return &books.ValidationReq{
		Book: &books.Book{
			Id:          id,
			Title:       title,
			Description: desc,
			Pages:       pages,
			Author:      author,
			Year:        year,
			Edition:     edition,
		},
	}
}
