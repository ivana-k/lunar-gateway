package server

import (
	"encoding/json"
	"fmt"
	pb "github.com/c12s/celestial/pb"
	"github.com/c12s/lunar-gateway/model"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read the request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := model.KVS{}
		if err := json.Unmarshal(body, &data); err != nil {
			sendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
			return
		}

		fmt.Println(data)
		fmt.Println(mutateToProto(data))

		//check rights

		//put to queue

		//return answer
		sendJSONResponse(w, map[string]string{"message": "success"})
	}
}

func mutateToProto(data model.KVS) *pb.MutateReq {
	labelsKV := []*pb.KV{}
	dataKV := []*pb.KV{}

	for _, item := range data.Labels {
		kv := &pb.KV{
			Key:   item.Key,
			Value: item.Value,
		}
		labelsKV = append(labelsKV, kv)
	}

	for _, item := range data.Data {
		kv := &pb.KV{
			Key:   item.Key,
			Value: item.Value,
		}
		dataKV = append(dataKV, kv)
	}

	labels := &pb.Label{
		Labels: labelsKV,
	}
	configs := &pb.Data{
		Data: dataKV,
	}

	req := &pb.MutateReq{
		RegionId:  data.RegionID,
		ClusterId: data.ClusterID,
		Label:     labels,
		Data:      configs,
		Kind:      pb.ReqKind_CONFIGS,
	}

	return req
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write the response body: %v", err)
		return
	}
}

func sendErrorMessage(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(status)
	io.WriteString(w, msg)
}
