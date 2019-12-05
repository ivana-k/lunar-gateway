package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func (server *LunarServer) setupTrace() {
	configs := server.r.PathPrefix("/trace").Subrouter()
	configs.HandleFunc("/list", auth(server.rightsList(server.listTraces()))).Methods("GET")
	configs.HandleFunc("/get", auth(server.rightsList(server.getTrace()))).Methods("GET")
}

func (s *LunarServer) listTraces() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.URL.Query()["tags"]; !ok {
			sendErrorMessage(w, errors.New("no tags in the request").Error(), http.StatusBadRequest)
			return
		}

		req := toListTrace(r.URL.Query()["tags"][0])
		client := NewStellarClient(s.clients[STELLAR])
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.List(ctx, req)
		if err != nil {
			sendErrorMessage(w, err.Error(), http.StatusBadRequest)
			return
		}

		rresp := traceListToJson(resp)
		data, rerr := json.Marshal(rresp)
		if rerr != nil {
			sendErrorMessage(w, rerr.Error(), http.StatusBadRequest)
			return
		}
		sendJSONResponse(w, string(data))
	}
}

func (s *LunarServer) getTrace() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.URL.Query()["traceId"]; !ok {
			sendErrorMessage(w, errors.New("no traceId in the request").Error(), http.StatusBadRequest)
			return
		}

		req := toGetTrace(r.URL.Query()["traceId"][0])
		client := NewStellarClient(s.clients[STELLAR])
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.Get(ctx, req)
		if err != nil {
			sendErrorMessage(w, err.Error(), http.StatusBadRequest)
			return
		}

		rresp := traceGetToJson(resp)
		data, rerr := json.Marshal(rresp)
		if rerr != nil {
			sendErrorMessage(w, rerr.Error(), http.StatusBadRequest)
			return
		}
		sendJSONResponse(w, string(data))
	}
}
