package handler_nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/otyang/yasante/pkg/config"
	"github.com/otyang/yasante/pkg/logger"
)

type Handler struct {
	Log    logger.Interface
	Config *config.Config
}

func NewHandler(config *config.Config, Log logger.Interface) *Handler {
	return &Handler{
		Log:    Log,
		Config: config,
	}
}

func (h *Handler) SubscribeFileUpload(msg *nats.Msg) {
	fmt.Println("STOCKS: " + string(msg.Subject) + " : " + string(msg.Data))
}
