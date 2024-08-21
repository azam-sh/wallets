package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	address string
	port    int
	mux     *mux.Router
}

func (s *Server) Run() {
	fmt.Println("Server running on", fmt.Sprintf("%v:%v", s.address, s.port))

	http.ListenAndServe(fmt.Sprintf("%v:%v", s.address, s.port), s.mux)
}

func NewServer(Address string, Port int, mux *mux.Router) *Server {
	return &Server{
		address: Address,
		port:    Port,
		mux:     mux,
	}
}
