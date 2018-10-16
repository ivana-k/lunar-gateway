package server

import (
	"encoding/json"
	"fmt"
	pb "github.com/c12s/celestial/pb"
	"github.com/c12s/lunar-gateway/model"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
)

func (server *LunarServer) setupConfigs() {
	configs := server.r.PathPrefix("/configs").Subrouter()

	configs.HandleFunc("/", server.getConfigs()).Methods("GET")
	configs.HandleFunc("/{regionid}", server.getRegionConfigs()).Methods("GET")
	configs.HandleFunc("/{regionid}/{clusterid}", server.getClusterConfigs()).Methods("GET")
	configs.HandleFunc("/new", server.createConfigs()).Methods("POST")
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
		//TODO: Check rights and so on...!!

		vars := mux.Vars(r)
		regionid := vars["regionid"]
		clusterid := vars["clusterid"]

		keys := r.URL.Query()
		var req pb.ListReq
		RequestToProto(keys, &req)
		req.RegionId = regionid
		req.ClusterId = clusterid

		resp, err := s.client.List(context.Background(), &req)
		if err != nil {
			sendErrorMessage(w, resp.Error, http.StatusBadRequest)
		}

		sendJSONResponse(w, "OK")
	}
}

func (s *LunarServer) createConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: Check rights and so on...!!!

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read the request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := model.MutateRequest{}
		if err := json.Unmarshal(body, &data); err != nil {
			sendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
			return
		}

		// var req pb.MutateRequest
		// RequestToProto(data, &req)
		// //Call celestiall RPC who should put job to queue and return answer sometinhg like job accepted!
		// resp, err := s.client.Mutate(context.Background(), &req)
		// if err != nil {
		// 	sendErrorMessage(w, "Error from Celestial Service!", http.StatusBadRequest)
		// }

		// if resp.Error == "NONE" {
		// 	sendErrorMessage(w, resp.Error, http.StatusBadRequest)
		// }

		//return answer
		sendJSONResponse(w, map[string]string{"message": "success"})
	}
}
