package server

import (
	aPb "github.com/c12s/scheme/apollo"
	bPb "github.com/c12s/scheme/blackhole"
	cPb "github.com/c12s/scheme/celestial"
	"google.golang.org/grpc"
	"log"
)

func NewCelestialClient(address string) cPb.CelestialServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to celestial service: %v", err)
	}

	return cPb.NewCelestialServiceClient(conn)
}

func NewBlackHoleClient(address string) bPb.BlackHoleServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to blackhole service: %v", err)
	}

	return bPb.NewBlackHoleServiceClient(conn)
}

func NewApolloClient(address string) aPb.ApolloServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to apollo service: %v", err)
	}

	return aPb.NewApolloServiceClient(conn)
}
