package server

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/c12s/blackhole/pb"
	"github.com/c12s/lunar-gateway/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

		data := model.NMutateRequest{}
		if err := json.Unmarshal(body, &data); err != nil {
			sendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
			return
		}

		var req *pb.PutReq
		RequestToProto(data, req)
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
