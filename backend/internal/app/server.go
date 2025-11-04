package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server encapsulates application dependencies. Add DB, logger, config here later.
type Server struct {
	addr string
	// add other fields like db, logger, config
}

// NewServer returns a new Server with default settings.
func NewServer() *Server {
	return &Server{
		addr: ":8080",
	}
}

// Run starts the HTTP server and blocks until it returns an error.
func (s *Server) Run() error {
	router := newRouter(s)

	// Use http.Server for future extensibility (timeouts, TLS, etc.)
	srv := &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	fmt.Printf("listening on %s\n", s.addr)
	return srv.ListenAndServe()
}

// Ensure Gin's engine type is imported even if not used directly in this file.
var _ = gin.New
