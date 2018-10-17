package server

import (
	"context"
	"encoding/json"
	"github.com/c12s/lunar-gateway/model"
	bPb "github.com/c12s/scheme/blackhole"
	cPb "github.com/c12s/scheme/celestial"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var caq = [...]string{"labels", "compare"}

func (server *LunarServer) setupConfigs() {
	configs := server.r.PathPrefix("/configs").Subrouter()
	configs.HandleFunc("/list", server.listConfigs()).Methods("GET")
	configs.HandleFunc("/mutate", server.mutateConfigs()).Methods("POST")
}

func (s *LunarServer) listConfigs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: Check rights and so on...!!!
		keys := r.URL.Query()

		var req *cPb.ListReq
		RequestToProto(keys, req)
		req.Kind = cPb.ReqKind_CONFIGS

		client := NewCelestialClient(s.clients[CELESTIAL])
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		cancel()

		resp, err := client.List(ctx, req)
		if err != nil {
			sendErrorMessage(w, resp.Error, http.StatusBadRequest)
		}

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

		var req *bPb.PutReq
		RequestToProto(data, req)
		req.Kind = bPb.TaskKind_CONFIGS

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
