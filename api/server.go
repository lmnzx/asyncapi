package api

import (
	"context"
	"sync"
	"time"

	"github.com/lmnzx/asyncapi/config"
	"github.com/lmnzx/asyncapi/store"

	"github.com/fasthttp/router"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

type Server struct {
	config *config.Config
	store  *store.Queries
}

func New(config *config.Config, store *store.Queries) *Server {
	return &Server{
		config: config,
		store:  store,
	}
}

func (s *Server) Start(ctx context.Context) {
	r := router.New()

	r.GET("/ping", handler(s.ping))
	r.POST("/signup", handler(s.signupHandler))

	server := &fasthttp.Server{
		Handler: NewLoggerMiddleware(r.Handler),
	}

	go func() {
		addr := s.config.GetAddr()
		log.Info().Str("addr", addr).Msg("server started")
		if err := server.ListenAndServe(addr); err != nil {
			log.Error().Err(err).Msg("failed to start server")
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		log.Info().Msg("shutdown signal received, stopping server")
		if err := server.ShutdownWithContext(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("failed to shutdown server")
		}
	}()

	wg.Wait()
}
