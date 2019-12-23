package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	stellar "github.com/c12s/stellar-go"
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
		span, _ := stellar.FromRequest(r, "listTraces")
		defer span.Finish()
		fmt.Println(span)

		if _, ok := r.URL.Query()["tags"]; !ok {
			span.AddLog(&stellar.KV{"query error", "no tags in the request"})
			sendErrorMessage(w, errors.New("no tags in the request").Error(), http.StatusBadRequest)
			return
		}

		req := toListTrace(r.URL.Query()["tags"][0])
		client := NewStellarClient(s.clients[STELLAR])
		ctx, cancel := context.WithTimeout(
			appendToken(
				stellar.NewTracedGRPCContext(nil, span),
				r.Header["Authorization"][0],
			),
			10*time.Second,
		)
		defer cancel()

		resp, err := client.List(ctx, req)
		if err != nil {
			span.AddLog(&stellar.KV{"stellar.list error", err.Error()})
			sendErrorMessage(w, err.Error(), http.StatusBadRequest)
			return
		}

		rresp := traceListToJson(resp)
		data, rerr := json.Marshal(rresp)
		if rerr != nil {
			span.AddLog(&stellar.KV{"proto to json error", rerr.Error()})
			sendErrorMessage(w, rerr.Error(), http.StatusBadRequest)
			return
		}
		sendJSONResponse(w, string(data))
	}
}

func (s *LunarServer) getTrace() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span, _ := stellar.FromRequest(r, "getTrace")
		defer span.Finish()
		fmt.Println(span)

		if _, ok := r.URL.Query()["traceId"]; !ok {
			span.AddLog(&stellar.KV{"queryr error", "no traceId in the request"})
			sendErrorMessage(w, errors.New("no traceId in the request").Error(), http.StatusBadRequest)
			return
		}

		req := toGetTrace(r.URL.Query()["traceId"][0])
		client := NewStellarClient(s.clients[STELLAR])
		ctx, cancel := context.WithTimeout(
			appendToken(
				stellar.NewTracedGRPCContext(nil, span),
				r.Header["Authorization"][0],
			),
			10*time.Second,
		)
		defer cancel()

		resp, err := client.Get(ctx, req)
		if err != nil {
			span.AddLog(&stellar.KV{"stellar.get error", err.Error()})
			sendErrorMessage(w, err.Error(), http.StatusBadRequest)
			return
		}

		rresp := traceGetToJson(resp)
		data, rerr := json.Marshal(rresp)
		if rerr != nil {
			span.AddLog(&stellar.KV{"proto to json error", rerr.Error()})
			sendErrorMessage(w, rerr.Error(), http.StatusBadRequest)
			return
		}
		sendJSONResponse(w, string(data))
	}
}
