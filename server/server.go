package server

import (
	celestialPb "github.com/c12s/celestial/pb"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type LunarServer struct {
	r       *mux.Router
	address string
	client  celestialPb.CelestialServiceClient

	//queue for createion of new configs, secrets, roles, ...
}

func NewServer(serverAddress, clientAddress string) *LunarServer {
	//create server struct
	server := &LunarServer{
		r:       mux.NewRouter(),
		address: serverAddress,
		client:  getRolesClient(clientAddress),
	}

	//setup routes
	server.setup()

	//if all is good return server
	return server
}

func (server *LunarServer) setup() {
	server.setupConfigs()
	server.setupSecrets()
}

func (server *LunarServer) Start() {
	http.ListenAndServe(server.address, server.r)
}

func getRolesClient(address string) celestialPb.CelestialServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}

	return celestialPb.NewCelestialServiceClient(conn)
}
