package server

import (
	celestialPb "github.com/c12s/celestial/pb"
	"github.com/c12s/lunar-gateway/model/configs"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	// "google.golang.org/grpc"
	// "log"
	"net/http"
	"os"
	"strings"
)

type LunarServer struct {
	r       *mux.Router
	address string
	client  celestialPb.CelestialServiceClient
}

func createRouter(version string) *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	prefix := strings.Join([]string{"/api", version}, "/")
	return r.PathPrefix(prefix).Subrouter()
}

func NewServer(conf *configs.Config) *LunarServer {
	//create server struct
	server := &LunarServer{
		r:       createRouter(conf.ConfVersion),
		address: conf.ServerConf.Address,
		// client:  getRolesClient(conf.ServicesConf.Celestial.Addr),
	}

	//setup routes
	server.setup()

	//if all is good return server
	return server
}

func (server *LunarServer) setup() {
	server.setupConfigs()
	server.setupSecrets()
	server.setupActions()
	server.setupNamespaces()
	// server.setupArtifacts()
}

func (server *LunarServer) Start() {
	// http.ListenAndServe(server.address, server.r)
	http.ListenAndServe(server.address, handlers.LoggingHandler(os.Stdout, server.r))
}

// func getRolesClient(address string) celestialPb.CelestialServiceClient {
// 	conn, err := grpc.Dial(address, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Failed to start gRPC connection: %v", err)
// 	}

// 	return celestialPb.NewCelestialServiceClient(conn)
// }
