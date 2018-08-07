package main

import (
	"github.com/gorilla/mux"
)

type LunarServer struct {
	r       *mux.Router
	address string
	port    string

	//queue for createion of new configs, secrets, roles, ...
}
