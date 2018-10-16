package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (server *LunarServer) setupNamespaces() {
	secrets := server.r.PathPrefix("/namespaces").Subrouter()
	secrets.HandleFunc("/list", server.listNamespaces()).Methods("GET")
	secrets.HandleFunc("/mutate", server.mutateNamespaces()).Methods("POST")
}

func (s *LunarServer) listNamespaces() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := r.URL.Query()
		fmt.Println(keys)
		sendJSONResponse(w, map[string]string{"status": "ok"})
	}
}

func (s *LunarServer) mutateNamespaces() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: Check rights and so on...!!!

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read the request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println(body)
		sendJSONResponse(w, map[string]string{"message": "success"})
	}
}
