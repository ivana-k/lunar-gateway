package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type LunarServer struct {
	r       *mux.Router
	address string
	port    string

	//queue for createion of new configs, secrets, roles, ...
}

func NewServer(address, port string) *LunarServer {
	//create server struct
	server := &LunarServer{
		mux.NewRouter(),
		address,
		port,
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

func (server *LunarServer) resolve() string {
	s := []string{server.address, server.port}
	return strings.Join(s, ":")
}

func (server *LunarServer) Start() {
	http.ListenAndServe(server.resolve(), server.r)
}
