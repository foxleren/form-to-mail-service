package models

import (
	"context"
	"github.com/siruspen/logrus"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		//TLSConfig:      &tls.Config{ServerName: os.Getenv("DOMAIN")},
	}
	logrus.Printf("Server is listening on %s port", port)

	return s.httpServer.ListenAndServe()
	//return s.httpServer.ListenAndServeTLS("certificate.crt", "private.key")
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
