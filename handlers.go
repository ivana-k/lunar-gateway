package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

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

func (server *LunarServer) Test() {
	fmt.Println("Test")
}

func (server *LunarServer) setup() {
	server.r.HandleFunc("/configs", server.getConfigs()).Methods("GET")
	server.r.HandleFunc("/configs/{regionid}", server.getRegionConfigs()).Methods("GET")
	server.r.HandleFunc("/configs/{regionid}/{clusterid}", server.getClusterConfigs()).Methods("GET")
	server.r.HandleFunc("/configs/{regionid}/{clusterid}/{nodeid}", server.getNodeConfigs()).Methods("GET")
	server.r.HandleFunc("/configs/{regionid}/{clusterid}/{nodeid}/{processid}", server.getProcessConfigs()).Methods("GET")

	server.r.HandleFunc("/configs/new", server.createConfigs()).Methods("POST")
}

func (s *LunarServer) getConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) getRegionConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) getClusterConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) getNodeConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) getProcessConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) createConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Create Configs")
	}
}

func (server *LunarServer) resolve() string {
	s := []string{server.address, server.port}

	return strings.Join(s, ":")
}

func (server *LunarServer) Start() {
	http.ListenAndServe(server.resolve(), server.r)
}
