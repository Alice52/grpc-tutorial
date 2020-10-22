package server

/**
Start our listener
Start a new gRPC server with no service registered yet.
Create a blog service and register it to our new gRPC server.
Connect to MongoDB.
Accept incoming connections on our listener
Handle user initiated server shutdown (CTRL+C)
*/
import (
	blog "../proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type BlogServiceServer struct {
}

func main() {
	// config log package
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50051...")

	// start listener. 50051 is the default gRPC port
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	// Set options, here we can configure things like TLS support
	opts := []grpc.serverOption{}
}
