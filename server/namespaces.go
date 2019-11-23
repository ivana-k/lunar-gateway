package server

import (
	"context"
	"encoding/json"
	"github.com/c12s/lunar-gateway/model"
	cPb "github.com/c12s/scheme/celestial"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (server *LunarServer) setupNamespaces() {
	secrets := server.r.PathPrefix("/namespaces").Subrouter()
	secrets.HandleFunc("/list", auth(server.rightsList(server.listNamespaces()))).Methods("GET")
	secrets.HandleFunc("/mutate", auth(server.rightsMutate(server.mutateNamespaces()))).Methods("POST")
}

var naq = [...]string{"user", "labels", "compare", "name"}

func (s *LunarServer) listNamespaces() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := listToProto(r.URL.Query(), cPb.ReqKind_NAMESPACES)
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

func (s *LunarServer) mutateNamespaces() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read the request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := &model.NMutateRequest{}
		if err := json.Unmarshal(body, data); err != nil {
			sendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
			return
		}

		req := mutateNSToProto(data)
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
