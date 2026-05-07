package transport

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	httpHandlers handlers
}

func New(awg awg, storagePath string) *server {
	return &server{
		httpHandlers: handlers{awg, storagePath},
	}
}

func (s *server) Start(endpoint string) {
	r := mux.NewRouter()
	r.HandleFunc("/peers/{publicKey}", s.httpHandlers.DeletePeer).Methods("DELETE")
	r.HandleFunc("/peers", s.httpHandlers.AddPeer).Methods("POST")
	r.HandleFunc("/peers/{id}/config", s.httpHandlers.SendConfFile).Methods("GET")

	fmt.Printf("HTTP started on %s\n", endpoint)
	http.ListenAndServe(endpoint, r)
}
