package server

import (
	"context"
	"encoding/json"
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

var caq = [...]string{"labels", "compare", "user"}

func (s *LunarServer) listConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: Check rights and so on...!!!

		keys := r.URL.Query()
		if _, ok := keys[user]; !ok {
			sendErrorMessage(w, "missing user id", http.StatusBadRequest)
			return
		}

		req := listToProto(keys)
		client := NewCelestialClient(s.clients[CELESTIAL])
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.List(ctx, req)
		if err != nil {
			sendErrorMessage(w, err.Error(), http.StatusBadRequest)
			return
		}

		rresp := protoToNSListResp(resp)
		data, rerr := json.Marshal(rresp)
		if rerr != nil {
			sendErrorMessage(w, rerr.Error(), http.StatusBadRequest)
			return
		}
		sendJSONResponse(w, string(data))
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

		data := &model.MutateRequest{}
		if err := json.Unmarshal(body, data); err != nil {
			sendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
			return
		}

		req := mutateToProto(data)
		client := NewBlackHoleClient(s.clients[BLACKHOLE])
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.Put(ctx, req)
		if err != nil {
			sendErrorMessage(w, "Error from Celestial Service!", http.StatusBadRequest)
			return
		}
		sendJSONResponse(w, map[string]string{"message": resp.Msg})
	}
}
