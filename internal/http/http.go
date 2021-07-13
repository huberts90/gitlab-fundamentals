package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const livenessPath = "/healthz"

type Server struct {
	logger *zap.Logger
	port   string
	router *mux.Router
}

func NewServer(logger *zap.Logger, port string) *Server {
	srv := &Server{
		logger: logger,
		port:   port,
		router: mux.NewRouter(),
	}

	// Register the routes
	srv.router.HandleFunc(livenessPath, srv.HandleLiveness)

	srv.logger.Info("Http configuration details",
		zap.String("PORT", srv.port),
		zap.String("livenessPath", livenessPath))

	return srv
}

func (server *Server) Serve() {
	server.logger.Info("Http server listening", zap.String("PORT", server.port))
	err := http.ListenAndServe(":"+server.port, server.router)
	if err != nil {
		server.logger.Fatal("Http server failed", zap.Error(err))
	}
}

// HandleHealthz is a liveness probe
func (server *Server) HandleLiveness(w http.ResponseWriter, _ *http.Request) {
	server.logger.Info("HTTP GET /healthz")
	w.WriteHeader(http.StatusOK)
}
