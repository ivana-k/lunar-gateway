package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/c12s/lunar-gateway/model"
	aPb "github.com/c12s/scheme/apollo"
	stellar "github.com/c12s/stellar-go"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (server *LunarServer) setupAuth() {
	auth := server.r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", server.login()).Methods("POST")
	auth.HandleFunc("/logout", server.logout()).Methods("GET")
	auth.HandleFunc("/register", server.register()).Methods("POST")
}

func (server *LunarServer) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trace := stellar.Init("lunar-gateway")
		defer trace.Finish()

		span := trace.Span("login")
		defer span.Finish()
		fmt.Println(span)
		span.AddTag(
			&stellar.KV{"app", "test"},
			&stellar.KV{"trace", "mytrace"},
		)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read the request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := &model.Credentials{}
		if err := json.Unmarshal(body, data); err != nil {
			sendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
			return
		}

		client := NewApolloClient(server.clients[APOLLO])
		ctx, cancel := context.WithTimeout(
			stellar.NewTracedGRPCContext(nil, span), 10*time.Second)
		defer cancel()

		req := &aPb.AuthOpt{
			Data: map[string]string{
				"intent":   "login",
				"username": data.Username,
				"password": data.Password,
			},
		}

		resp, err := client.Auth(ctx, req)
		if err != nil {
			sendErrorMessage(w, err.Error(), http.StatusBadRequest)
			return
		}

		if resp.Value {
			h := map[string]string{
				"Content-Type": "application/json; charset=UTF-8",
				"Auth-Token":   resp.Data["token"],
			}
			sendJSONResponseWithHeader(w, map[string]string{"message": resp.Data["message"]}, h)
			return
		}

		sendJSONResponse(w, map[string]string{"message": resp.Data["message"]})
	}
}

func (server *LunarServer) logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendJSONResponse(w, map[string]string{"message": "logout"})
	}
}

func (server *LunarServer) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trace := stellar.Init("lunar-gateway")
		defer trace.Finish()

		span := trace.Span("register")
		defer span.Finish()
		fmt.Println(span)
		span.AddTag(
			&stellar.KV{"app", "test"},
			&stellar.KV{"trace", "mytrace"},
		)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read the request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := &model.UMutateRequest{}
		if err := json.Unmarshal(body, data); err != nil {
			sendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
			return
		}

		client := NewApolloClient(server.clients[APOLLO])
		ctx, cancel := context.WithTimeout(
			appendToken(
				stellar.NewTracedGRPCContext(nil, span),
				"my_test_token", //TODO: this shoulbd be created prior to register
			), 10*time.Second)
		defer cancel()

		req := &aPb.AuthOpt{
			Data: map[string]string{
				"intent":       "register",
				"username":     data.Info["username"],
				"password":     data.Info["password"],
				"firstname":    data.Info["firstname"],
				"lastname":     data.Info["lastname"],
				"organisation": data.Info["organisation"],
				"role":         data.Info["role"],
			},
			Extras: map[string]*aPb.OptExtras{},
		}

		for lk, lv := range data.Labels {
			req.Extras[lk] = &aPb.OptExtras{Data: []string{lv}}
		}

		resp, err := client.Auth(ctx, req)
		if err != nil {
			sendErrorMessage(w, err.Error(), http.StatusBadRequest)
			return
		}

		if resp.Value {
			h := map[string]string{
				"Content-Type": "application/json; charset=UTF-8",
				"Auth-Token":   resp.Data["token"],
			}
			sendJSONResponseWithHeader(w, map[string]string{"message": resp.Data["message"]}, h)
			return
		}
		sendJSONResponse(w, map[string]string{"message": resp.Data["message"]})
	}
}
