package server

import (
	"fmt"
	"net/http"
)

func (server *LunarServer) setupConfigs() {
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
