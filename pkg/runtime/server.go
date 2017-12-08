package runtime

import (
	"fmt"
	"net/http"
)

type Server interface {
	Run() error
}

type serverImpl struct {
	mux  *http.ServeMux
	port int
}

func (s *serverImpl) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}

func NewServer(port int) Server {
	// TODO: create the mux
	return &serverImpl{port: port}
}
