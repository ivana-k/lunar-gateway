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

var aaq = [...]string{"labels", "compare", "from", "to", "top"}

func (server *LunarServer) setupActions() {
	secrets := server.r.PathPrefix("/actions").Subrouter()
	secrets.HandleFunc("/list", server.listSecrets()).Methods("GET")
	secrets.HandleFunc("/mutate", server.mutateActions()).Methods("POST")
}

func (s *LunarServer) listActions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: Check rights and so on...!!!
		keys := r.URL.Query()
		extras := []*cPb.KV{}
		if val, ok := keys[user]; ok {
			extras = append(extras, &cPb.KV{Key: user, Value: val[0]})
		} else {
			sendErrorMessage(w, "missing user id", http.StatusBadRequest)
		}

		var req *cPb.ListReq
		RequestToProto(keys, req)
		req.Kind = cPb.ReqKind_ACTIONS
		req.Extras = append(req.Extras, extras...)

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

func (s *LunarServer) mutateActions() http.HandlerFunc {
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
		req.Kind = bPb.TaskKind_ACTIONS

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
