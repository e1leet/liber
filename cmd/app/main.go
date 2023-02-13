package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/e1leet/liber/internal/app"
	"github.com/e1leet/liber/internal/config"
	"github.com/e1leet/liber/internal/utils/common"
	"github.com/rs/zerolog/log"

	_ "github.com/e1leet/liber/docs"
)

//	@title			Liber
//	@version		0.0.0
//	@description	Liber API documentation

//	@contact.name	Damir mirasov
//	@contact.url	https://github.com/e1leet
//	@contact.email	damirmirasovmain@gmail.com

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit-0/

//	@host		localhost:8000
//	@BasePath	/api

func main() {
	cfgPath := common.ConfigPath()

	cfg, err := config.New(cfgPath)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	common.ConfigureLogging(cfg.Log.Level)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	a := app.New(cfg, log.Logger)
	if err := a.Run(ctx); err != nil {
		stop()
		log.Fatal().Err(err).Send() //nolint:gocritic
	}
}
