package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	HTTPServer *http.Server
}

func (s *Server) Start() error {
	return s.HTTPServer.ListenAndServe()
}

func NewServer(router *gin.Engine) *Server {
	server := &http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	return &Server{
		HTTPServer: server,
	}
}
