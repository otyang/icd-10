package icd

import (
	"context"

	"github.com/gofiber/fiber/v2"

	eventDTOs "github.com/otyang/icd-10/internal/event"
	"github.com/otyang/icd-10/internal/icd/entity"
	handlerEvent "github.com/otyang/icd-10/internal/icd/handler_event"
	handlerHttp "github.com/otyang/icd-10/internal/icd/handler_http"
	"github.com/otyang/icd-10/internal/icd/repository/bun"

	"github.com/otyang/icd-10/pkg/config"
	"github.com/otyang/icd-10/pkg/datastore"
	loggers "github.com/otyang/icd-10/pkg/logger"
	mw "github.com/otyang/icd-10/pkg/middleware"

	"github.com/gookit/event"
)

func RegisterHttpHandlers(
	ctx context.Context, router *fiber.App,
	config *config.Config, log loggers.Interface, db datastore.OrmDB,
) {
	var (
		repo    = bun.NewICDRepository(db)
		handler = handlerHttp.NewHandler(repo, config, log)
	)

	// Using default router
	{
		router.Get("/", handler.Welcome)
		router.Get("/icd/", handler.List)
		router.Get("/icd/:fullCode", handler.Get)
		router.Delete("/icd/:fullCode", handler.Delete)
		router.Post("/icd-upload", handler.Upload)

		router.Post("/icd", mw.ValidateBody[entity.CreateICDRequest], handler.Create)
		router.Put("/icd/:fullCode", mw.ValidateBody[entity.EditICDRequest], handler.Edit)
	}
}

func RegisterEventsHandlers(ctx context.Context, config *config.Config, log loggers.Interface) {

	handler := handlerEvent.NewHandler(config, log)

	event.On(
		eventDTOs.SubjectFileUpload,
		event.ListenerFunc(handler.EventHandlerFileUpload),
		event.Normal,
	)
}
