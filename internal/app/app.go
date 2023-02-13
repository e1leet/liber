package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/e1leet/liber/internal/config"
	"github.com/e1leet/liber/internal/transport/middleware"
	"github.com/e1leet/liber/pkg/errors"
	"github.com/e1leet/liber/pkg/shutdown"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	router chi.Router
	cfg    *config.Config
	closer *shutdown.Closer
	logger zerolog.Logger
}

func New(cfg *config.Config, logger zerolog.Logger) *App {
	return &App{
		cfg:    cfg,
		router: chi.NewRouter(),
		closer: &shutdown.Closer{},
		logger: logger,
	}
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Info().Msg("configure dependencies")

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.cfg.Server.Host, a.cfg.Server.Port),
		Handler: a.router,
	}

	a.logger.Info().Msg("configure middlewares")

	a.logger.Debug().Msg("configure CORS")
	a.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   a.cfg.CORS.AllowedOrigins,
		AllowedMethods:   a.cfg.CORS.AllowedMethods,
		AllowedHeaders:   a.cfg.CORS.AllowedHeaders,
		ExposedHeaders:   a.cfg.CORS.ExposedHeaders,
		MaxAge:           a.cfg.CORS.MaxAge,
		AllowCredentials: a.cfg.CORS.AllowCredentials,
	}))

	a.router.Use(chiMiddleware.AllowContentType("application/json"))
	a.router.Use(middleware.LoggerMiddleware(a.logger))

	a.logger.Info().Msg("configure controllers")

	a.logger.Debug().Msg("configure swagger controller")
	a.router.Get("/swagger/*", httpSwagger.WrapHandler)
	a.router.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})

	a.logger.Info().Msg("configure closer")
	a.closer.Add(srv.Shutdown)

	// Run server
	go func() {
		a.logger.Info().Msgf("running server on %s", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal().Err(err).Send()
		}
	}()

	<-ctx.Done()

	a.logger.Info().Msg("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		a.cfg.Server.ShutdownTimeout,
	)
	defer cancel()

	if err := a.closer.Close(shutdownCtx); err != nil {
		return errors.Wrap(err, "failed to shutdown")
	}

	return nil
}
