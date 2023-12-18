package startup

import (
	"context"
	"fmt"
	"gateway/client"
	"gateway/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{config: config}
}

func (s *Server) Start() {
	clientRegistry := s.prepareClients()
	router := s.prepareRoutes(clientRegistry)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.config.Gateway.Port), router))
}

func (s *Server) prepareClients() *client.ClientRegistry {
	log.Println("Preparing clients")
	clientRegistry := &client.ClientRegistry{
		Clients: make(map[string]client.Client),
	}

	for k, v := range s.config.Services {
		clientRegistry.NewClient(k, v)
	}

	return clientRegistry
}
func (s *Server) prepareRoutes(clientRegistry *client.ClientRegistry) *mux.Router {
	router := mux.NewRouter().PathPrefix(s.config.Gateway.Route).Subrouter()

	for group, versions := range s.config.Groups {
		groupRouter := router.PathPrefix("/" + group).Subrouter()
		for version, methods := range versions {
			versionRouter := groupRouter.PathPrefix("/" + version).Subrouter()
			for mtdName, mtdConf := range methods {
				log.Printf("Name %s Conf %+v", mtdName, mtdConf)
				client := clientRegistry.Clients[mtdConf.Service]
				versionRouter.Path(mtdConf.MethodRoute).HandlerFunc(methodNameMiddleware(mtdName, client.InvokeGrpcMethod)).Methods(mtdConf.Type)
			}
		}
	}

	return router
}

func methodNameMiddleware(mtdName string, h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "mtdName", mtdName)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
