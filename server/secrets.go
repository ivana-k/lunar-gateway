package server

import (
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (server *LunarServer) setupSecrets() {
	secrets := server.r.PathPrefix("/secrets").Subrouter()

	secrets.HandleFunc("/", server.getSecrets()).Methods("GET")
	secrets.HandleFunc("/{regionid}", server.getRegionSecrets()).Methods("GET")
	secrets.HandleFunc("/{regionid}/{clusterid}", server.getClusterSecrets()).Methods("GET")
	secrets.HandleFunc("/new", server.createSecrets()).Methods("POST")
}

func (s *LunarServer) getSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get Configs")
	}
}

func (s *LunarServer) getRegionSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		regionid := vars["regionid"]

		fmt.Fprintf(w, "Get Configs region:%s", regionid)
	}
}

func (s *LunarServer) getClusterSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		regionid := vars["regionid"]
		clusterid := vars["clusterid"]

		fmt.Fprintf(w, "Get Configs region:%s cluster:%s", regionid, clusterid)
	}
}

func (s *LunarServer) createSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check rights

		//put to queue

		//return answer
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Create Configs")
	}
}
