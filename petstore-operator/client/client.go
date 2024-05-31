package client

import (
	pb "github.com/PacktPublishing/Go-for-DevOps/chapter/14/petstore-operator/client/proto"
	"google.golang.org/grpc"
)

// Client is a client to the petstore service.
type Client struct {
	client pb.PetStoreClient
	conn   *grpc.ClientConn
}
