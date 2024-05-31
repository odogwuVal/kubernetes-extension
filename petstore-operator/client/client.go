package client

import (
	"context"
	"time"

	"github.com/PacktPublishing/Go-for-DevOps/chapter/14/petstore-operator/client/internal/server/storage"
	pb "github.com/PacktPublishing/Go-for-DevOps/chapter/14/petstore-operator/client/proto"
	"google.golang.org/grpc"
)

// Client is a client to the petstore service.
type Client struct {
	client pb.PetStoreClient
	conn   *grpc.ClientConn
}

// New is the constructor for Client. addr is the server's [host]:[port].
func New(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		client: pb.NewPetStoreClient(conn),
		conn:   conn,
	}, nil
}

// Pet is a wrapper around a *pb.Pet that can return Go versions of
// fields and errors if the returned stream has an error.
type Pet struct {
	*pb.Pet
	err error
}

// Proto will give the Pet's proto representation.
func (p Pet) Proto() *pb.Pet {
	return p.Pet
}

// Birthday returns the Pet's birthday as a time.Time.
func (p Pet) Birthday() time.Time {
	// We are ignoring the error as we will either get a zero time
	// anyways and the server should be preventing this problem.
	t, _ := storage.BirthdayToTime(context.Background(), p.Pet.Birthday)
	return t
}

// Error indicates if there was an error in the Pet output stream.
func (p Pet) Error() error {
	return p.err
}

// CallOptions are optional options for an RPC call.
type CallOption func(co *callOptions)

type callOptions struct {
	trace *string
}

// TraceID will cause the RPC call to execute a trace on the service and return "s" to the ID.
// If s == nil, this will ignore the option. If "s" is not set after the call finishes, then
// no trace was made.
func TraceID(s *string) CallOption {
	return func(co *callOptions) {
		if s == nil {
			return
		}
		co.trace = s
	}
}
