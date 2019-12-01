package server

import (
	"fmt"
	"github.com/c12s/lunar-gateway/model/configs"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

type LunarServer struct {
	r       *mux.Router
	address string
	clients map[string]string
}

func createBaseRouter(version string) *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	prefix := strings.Join([]string{"/api", version}, "/")
	return r.PathPrefix(prefix).Subrouter()
}

func NewServer(conf *configs.Config) *LunarServer {
	//create server struct
	server := &LunarServer{
		r:       createBaseRouter(conf.ConfVersion),
		address: conf.ServerConf.Address,
		clients: conf.ServicesConf,
	}

	//setup routes
	server.setupEndpoints()

	//if all is good return server
	return server
}

func (server *LunarServer) setupEndpoints() {
	server.setupConfigs()
	server.setupSecrets()
	server.setupActions()
	server.setupNamespaces()
	server.setupAuth()
	server.setupTrace()
}

func (server *LunarServer) Start() {
	fmt.Println("LunarServer Started")
	http.ListenAndServe(server.address, handlers.LoggingHandler(os.Stdout, server.r))
}
