package server

import (
	"net/http"
)

func (server *LunarServer) setupAuth() {
	auth := server.r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", server.login()).Methods("POST")
	auth.HandleFunc("/logout", server.logout()).Methods("GET")
}

func (server *LunarServer) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendJSONResponse(w, map[string]string{"message": "logged in"})
	}
}

func (server *LunarServer) logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendJSONResponse(w, map[string]string{"message": "logout"})
	}
}
