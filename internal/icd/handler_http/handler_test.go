package handler_http

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/otyang/icd-10/internal/icd/entity"
	"github.com/otyang/icd-10/pkg/config"
	"github.com/otyang/icd-10/pkg/logger"
)

func TestHandler_Get(t *testing.T) {
	type fields struct {
		Log    logger.Interface
		Config *config.Config
		Repo   entity.IICDRepository
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Log:    tt.fields.Log,
				Config: tt.fields.Config,
				Repo:   tt.fields.Repo,
			}
			if err := h.Get(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
