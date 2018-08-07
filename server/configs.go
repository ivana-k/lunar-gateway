package server

import (
	"encoding/json"
	"fmt"
	"github.com/c12s/lunar-gateway/model"
	"github.com/gorilla/mux"
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
		vars := mux.Vars(r)
		regionid := vars["regionid"]

		fmt.Fprintf(w, "Get Configs region:%s", regionid)
	}
}

func (s *LunarServer) getClusterConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		regionid := vars["regionid"]
		clusterid := vars["clusterid"]

		fmt.Fprintf(w, "Get Configs region:%s, cluster:%s", regionid, clusterid)
	}
}

func (s *LunarServer) getNodeConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		regionid := vars["regionid"]
		clusterid := vars["clusterid"]
		nodeid := vars["nodeid"]

		fmt.Fprintf(w, "Get Configs region:%s, cluster:%s, node:%s", regionid, clusterid, nodeid)
	}
}

func (s *LunarServer) getProcessConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		regionid := vars["regionid"]
		clusterid := vars["clusterid"]
		nodeid := vars["nodeid"]
		processid := vars["processid"]

		fmt.Fprintf(w, "Get Configs region:%s, cluster:%s, node:%s, process:%s", regionid, clusterid, nodeid, processid)
	}
}

func (s *LunarServer) createConfigs() http.HandlerFunc {
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
