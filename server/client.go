package server

import (
	bPb "github.com/c12s/blackhole/pb"
	cPb "github.com/c12s/celestial/pb"
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
