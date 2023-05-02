package main

import (
	"context"
	"flag"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/otyang/icd-10/internal/icd"
	"github.com/otyang/icd-10/pkg/config"
	"github.com/otyang/icd-10/pkg/datastore"
	"github.com/otyang/icd-10/pkg/logger"
	"github.com/otyang/icd-10/pkg/middleware"
)

func main() {
	var (
		configFile = flag.String("configFile", ".example.env", "full path to config file")

		cfg = func() *config.Config {
			flag.Parse()
			c := &config.Config{}
			config.Load(*configFile, c)
			return c
		}()
	)

	var (
		router      = fiber.New(fiber.Config{})
		ctx, cancel = context.WithCancel(context.Background())
		zlog        = logger.NewSlogLogger("debug", "json", os.Stdout)
		db          = datastore.NewDBConnection(cfg.Database.Driver, cfg.Database.URL, cfg.Database.PoolMax, cfg.Database.PrintQueriesToStdout)
	)

	logger.WithBaseInfo(zlog, cfg.AppName, cfg.AppAddress)

	defer cancel()
	defer func() {
		if err := db.Close(); err != nil {
			zlog.Fatal("Error closing database: " + err.Error())
		}
	}()

	{
		router.Use(
			middleware.Cors(),
			middleware.HttpLog(zlog),
		)

		icd.RegisterHttpHandlers(ctx, router, cfg, zlog, db)
		icd.RegisterEventsHandlers(ctx, cfg, zlog)
	}

	if err := router.Listen(cfg.AppAddress); err != nil {
		zlog.Fatal("Error starting server: " + err.Error())
	}
}
