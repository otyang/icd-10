package main

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/otyang/icd-10/internal/icd"
	"github.com/otyang/icd-10/pkg/config"
	"github.com/otyang/icd-10/pkg/datastore"
	"github.com/otyang/icd-10/pkg/logger"
	"github.com/otyang/icd-10/pkg/middleware"
)

var configFile = ""

func main() {
	cfg := &config.Config{}
	config.Load(configFile, cfg)

	var (
		router      = fiber.New(fiber.Config{})
		ctx, cancel = context.WithCancel(context.Background())
		zlog        = logger.NewSlogLogger("debug", "json", os.Stdout)
		db          = datastore.NewDBConnection(cfg.DBDriver, cfg.DBURL, cfg.DBPoolMax, cfg.DBPrintQueriesToStdout)
	)

	logger.WithBaseInfo(zlog, cfg.AppName, cfg.AppAddress)

	// pubSub := datastore.NewNatsFromCredential(cfg.NatsURL, cfg.NatsCredentialFile, zlog)

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

		// nats above http server, cause http blocking server.
		icd.RegisterHttpHandlers(ctx, router, cfg, zlog, db)
		// icd10.RegisterEventsHandlers(ctx, pubSub, cfg, zlog, db)
	}

	if err := router.Listen(cfg.AppAddress); err != nil {
		zlog.Fatal("Error starting server: " + err.Error())
	}
}

func createFolderForUpload(dir string, log logger.Interface) error {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		log.Fatal("upload directory creation error" + err.Error())
	}
	return nil
}
