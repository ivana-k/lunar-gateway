package server

import (
	celestialPb "github.com/c12s/celestial/pb"
	"github.com/c12s/lunar-gateway/model/configs"
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

func NewServer(conf *configs.Config) *LunarServer {
	//create server struct
	server := &LunarServer{
		r:       mux.NewRouter(),
		address: conf.ServerConf.Address,
		client:  getRolesClient(conf.ServicesConf.Celestial.Addr),
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
