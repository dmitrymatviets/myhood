package server

import (
	"context"
	"github.com/dmitrymatviets/myhood/infrastructure/logger"
	"net"
	"net/http"
	"strconv"

	"github.com/dmitrymatviets/myhood/infrastructure/server/config"
	"github.com/dmitrymatviets/myhood/infrastructure/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type HTTPServer struct {
	Ctx    context.Context
	cancel context.CancelFunc
	cfg    *config.ServerConfig
	engine *gin.Engine
	logger *logger.Logger
	server *http.Server
}

func NewHTTPServer(cfg config.ServerConfig, logger *logger.Logger) *HTTPServer {
	engine := gin.New()

	engine.Use(middleware.RecoveryMiddleware(&cfg, logger))
	engine.Use(middleware.RequestMiddleware(logger))
	engine.Use(middleware.ResponseMiddleware(&cfg, logger))
	engine.NoRoute(middleware.NoRoute())

	srv := &http.Server{
		Addr:         net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Handler:      engine,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &HTTPServer{
		Ctx:    ctx,
		cancel: cancel,
		cfg:    &cfg,
		engine: engine,
		server: srv,
		logger: logger,
	}
}

func (s *HTTPServer) AddRoutes(routes ...*Route) {
	for _, route := range routes {
		s.engine.Handle(route.Method, route.Path, route.HandleFuncs...)
	}
}

func (s *HTTPServer) Start() {
	go func() {
		s.logger.Info(context.Background(), "server started")
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.logger.Fatal(context.Background(), err.Error())
			return
		}
	}()
}

func (s *HTTPServer) Stop() error {
	defer s.logger.Info(context.Background(), "server stopped")
	s.cancel()

	var ctx context.Context
	if s.cfg.ShutdownTimeout != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
		defer cancel()
	} else {
		ctx = context.Background()
	}

	err := s.server.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, "shutdown failed")
	}
	return nil
}
