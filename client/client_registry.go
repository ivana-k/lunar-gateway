package client

import (
	"context"
	"time"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientRegistry struct {
	Clients map[string]Client
}

func (registry *ClientRegistry) NewClient(name, address string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var cc *grpc.ClientConn
	var err error
	defer cancel()
	if cc, err = grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock()); err != nil {
		panic(err)
	}
	refClient := grpcreflect.NewClientAuto(context.Background(), cc)
	descSource := grpcurl.DescriptorSourceFromServer(context.Background(), refClient)
	services, _ := descSource.ListServices()
	registry.Clients[name] = Client{refClient: refClient, cc: cc, descSource: descSource, grpcServiceName: services[0]}
}
