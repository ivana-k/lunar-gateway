package server

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/c12s/blackhole/pb"
	"github.com/c12s/lunar-gateway/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (server *LunarServer) setupConfigs() {
	configs := server.r.PathPrefix("/configs").Subrouter()
	configs.HandleFunc("/list", server.listConfigs()).Methods("GET")
	configs.HandleFunc("/mutate", server.mutateConfigs()).Methods("POST")
}

func (s *LunarServer) listConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := r.URL.Query()
		fmt.Println(keys)
		sendJSONResponse(w, map[string]string{"status": "ok"})
	}
}

func (s *LunarServer) mutateConfigs() http.HandlerFunc {
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

		var req *pb.PutReq
		RequestToProto(data, req)
		client := NewBlackHoleClient(s.clients[BLACKHOLE])
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		cancel()

		resp, err := client.Put(ctx, req)
		if err != nil {
			sendErrorMessage(w, "Error from Celestial Service!", http.StatusBadRequest)
		}

		sendJSONResponse(w, map[string]string{"message": resp.Msg})
	}
}
