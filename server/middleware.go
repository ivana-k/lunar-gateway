package server

import (
	"context"
	"encoding/json"
	"github.com/c12s/lunar-gateway/model"
	aPb "github.com/c12s/scheme/apollo"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Header["Authentication"]; !ok {
			sendErrorMessage(w, "missing authorization token", http.StatusBadRequest)
			return
		}

		if _, ok := r.URL.Query()["user"]; !ok {
			sendErrorMessage(w, "missing user id", http.StatusBadRequest)
			return
		}
		next(w, r)
	})
}

func (server *LunarServer) rightsList(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query()) == 0 {
			sendErrorMessage(w, "missing query parameters", http.StatusBadRequest)
			return
		}
		extras := map[string]*aPb.OptExtras{}
		for k, v := range r.URL.Query() {
			extras[k] = &aPb.OptExtras{Data: v}
		}

		spl := strings.Split(r.URL.Path, "/")
		req := &aPb.AuthOpt{
			Data: map[string]string{
				"intent": spl[4],
				"kind":   spl[3],
				"user":   r.URL.Query()["user"][0],
				"token":  r.Header["Authorization"][0],
			},
			Extras: extras,
		}

		client := NewApolloClient(server.clients[APOLLO])
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.Auth(ctx, req)
		if err != nil {
			sendErrorMessage(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !resp.Value {
			sendErrorMessage(w, resp.Data["message"], http.StatusBadRequest)
			return
		}
		next(w, r)
	})
}

func (server *LunarServer) rightsMutate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		spl := strings.Split(r.URL.Path, "/")
		extras := map[string]*aPb.OptExtras{}
		req := &aPb.AuthOpt{}
		if spl[3] != "namespaces" {
			data := &model.MutateRequest{}
			if err := json.Unmarshal(body, data); err != nil {
				sendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
				return
			}

			for _, r := range data.Regions {
				tmp := []string{}
				for _, c := range r.Clusters {
					tmp = append(tmp, c.ID)
				}
				extras[r.ID] = &aPb.OptExtras{Data: tmp}
			}

			req.Data = map[string]string{
				"intent":       spl[4],
				"kind":         spl[3],
				"user":         r.URL.Query()["user"][0],
				"token":        r.Header["Authorization"][0],
				"namespace":    data.MTData.Namespace,
				"queue":        data.MTData.Queue,
				"forceNSQueue": strconv.FormatBool(data.MTData.ForceNSQueue),
			}
			req.Extras = extras
		} else {
			data := &model.NMutateRequest{}
			if err := json.Unmarshal(body, data); err != nil {
				log.Println(err)
				sendErrorMessage(w, "Could not decode the request body as JSON", http.StatusBadRequest)
				return
			}

			req.Data = map[string]string{
				"intent":       spl[4],
				"kind":         spl[3],
				"user":         r.URL.Query()["user"][0],
				"token":        r.Header["Authorization"][0],
				"namespace":    data.MTData.Namespace,
				"queue":        data.MTData.Queue,
				"forceNSQueue": strconv.FormatBool(data.MTData.ForceNSQueue),
			}
			req.Extras = extras
		}

		client := NewApolloClient(server.clients[APOLLO])
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.Auth(ctx, req)
		if err != nil {
			log.Println(err)
			sendErrorMessage(w, "Error from Apollo Service!", http.StatusBadRequest)
			return
		}

		if !resp.Value {
			sendErrorMessage(w, resp.Data["message"], http.StatusBadRequest)
			return
		}
		next(w, r)
	})
}
