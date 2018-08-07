package server

import (
	"fmt"
	"net/http"
)

func (server *LunarServer) setupSecrets() {
	server.r.HandleFunc("/secrets", server.getSecrets()).Methods("GET")
	server.r.HandleFunc("/secrets/{regionid}", server.getRegionSecrets()).Methods("GET")
	server.r.HandleFunc("/secrets/{regionid}/{clusterid}", server.getClusterSecrets()).Methods("GET")
	server.r.HandleFunc("/secrets/{regionid}/{clusterid}/{nodeid}", server.getNodeSecrets()).Methods("GET")
	server.r.HandleFunc("/secrets/{regionid}/{clusterid}/{nodeid}/{processid}", server.getProcessSecrets()).Methods("GET")

	server.r.HandleFunc("/secrets/new", server.createSecrets()).Methods("POST")
}

func (s *LunarServer) getSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) getRegionSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) getClusterSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) getNodeSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) getProcessSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) createSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Create Configs")
	}
}
