package server

import (
	"encoding/json"
	"fmt"
	"github.com/c12s/lunar-gateway/model"
	"github.com/gorilla/mux"
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

func (s *LunarServer) getNodeSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		regionid := vars["regionid"]
		clusterid := vars["clusterid"]
		nodeid := vars["nodeid"]

		fmt.Fprintf(w, "Get Configs region:%s, cluster:%s, nodeid:%s", regionid, clusterid, nodeid)
	}
}

func (s *LunarServer) getProcessSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		regionid := vars["regionid"]
		clusterid := vars["clusterid"]
		nodeid := vars["nodeid"]
		processid := vars["processid"]

		fmt.Fprintf(w, "Get Configs region:%s, cluster:%s, node:%s, process:%s", regionid, clusterid, nodeid, processid)
	}
}

func (s *LunarServer) createSecrets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := model.Data{}

		//get data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		//check rights

		//put to queue

		//return answer
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Create Configs")
	}
}
