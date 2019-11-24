package server

import (
	"context"
	"encoding/json"
	"github.com/c12s/lunar-gateway/model"
	aPb "github.com/c12s/scheme/apollo"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (server *LunarServer) setupAuth() {
	auth := server.r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", server.login()).Methods("POST")
	auth.HandleFunc("/logout", server.logout()).Methods("GET")
}

func (server *LunarServer) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
			sendErrorMessage(w, "Error from Apollo Service!", http.StatusBadRequest)
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
