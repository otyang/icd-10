package icd

import (
	"context"

	"github.com/gofiber/fiber/v2"

	entity "github.com/otyang/icd-10/internal/icd/entity"
	handlerHttp "github.com/otyang/icd-10/internal/icd/handler_http"
	"github.com/otyang/icd-10/internal/icd/repository/bun"

	"github.com/otyang/icd-10/pkg/config"
	"github.com/otyang/icd-10/pkg/datastore"
	loggers "github.com/otyang/icd-10/pkg/logger"
	mw "github.com/otyang/icd-10/pkg/middleware"
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

// func RegisterEventsHandlers(
// 	ctx context.Context,
// 	pubSub *datastore.PubSub,
// 	config *config.Config,
// 	log loggers.Interface,
// 	db datastore.OrmDB,
// ) {
// 	var (
// 		repo    = bun.NewICDRepository(db)
// 		handler = handlerNats.NewHandler(repo, config, log)
// 	)

// 	{
// 		pubSub.Subscribe(context.TODO(), event.SubjectStockUpdates, handler.SubscribeStocks)
// 		pubSub.Subscribe(context.TODO(), event.SubjectSMSRecieve, handler.SubscribeSMSRecieve)

// 		_ = pubSub.Publish(context.TODO(), event.SubjectSMSRecieve, &event.Stock{
// 			Symbol: "GOOG",
// 			Price:  200,
// 		})

// 		err := pubSub.Publish(context.TODO(), event.SubjectStockUpdates, &event.SMSRecieve{
// 			From:    "+9999999999",
// 			Message: "Hello World!",
// 		})
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 	}
// }
