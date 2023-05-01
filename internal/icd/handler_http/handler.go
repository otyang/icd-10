package handler_http

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/otyang/icd-10/internal/icd/entity"
	"github.com/otyang/icd-10/pkg/config"
	"github.com/otyang/icd-10/pkg/logger"
	"github.com/otyang/icd-10/pkg/middleware"
	"github.com/otyang/icd-10/pkg/response"

	"github.com/gookit/event"
	eventDTOs "github.com/otyang/icd-10/internal/event"
)

type Handler struct {
	Log    logger.Interface
	Config *config.Config
	Repo   entity.IICDRepository
}

func NewHandler(repo entity.IICDRepository, config *config.Config, Log logger.Interface) *Handler {
	return &Handler{
		Log:    Log,
		Config: config,
		Repo:   repo,
	}
}

func (h *Handler) Welcome(c *fiber.Ctx) error {

	resp := response.Ok("", "Hello, welcome to the icd_10 page")
	return c.
		Status(resp.StatusCode).
		JSON(resp)
}

func (h *Handler) Get(c *fiber.Ctx) error {
	icdFullCode := c.Params("fullCode")

	result, err := h.Repo.GetRecord(context.TODO(), icdFullCode)
	if err != nil {
		if entity.IsErrNotFound(err) {
			return response.
				NotFound(
					"icd record not found", nil,
				)
		}
		return response.InternalServerError(err.Error(), nil)
	}

	resp := response.Ok("", result)
	return c.Status(resp.StatusCode).JSON(resp)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	body, ok := middleware.ValidatedDataFromContext(c).(*entity.CreateICDRequest)
	if !ok {
		return response.InternalServerError("error from ur end, invalid context", nil)
	}

	_, err := h.Repo.GetRecord(context.TODO(), body.GetFullCode())
	if err != nil && !entity.IsErrNotFound(err) {
		return response.InternalServerError(err.Error(), nil)
	}

	if err == nil {
		return response.Conflict("icd code already exists", nil)
	}

	result, err := h.Repo.CreateRecord(context.TODO(), *body)
	if err != nil {
		return response.InternalServerError(err.Error(), nil)
	}

	resp := response.Created("", result)
	return c.Status(resp.StatusCode).JSON(resp)
}

func (h *Handler) Edit(c *fiber.Ctx) error {
	body, ok := middleware.ValidatedDataFromContext(c).(*entity.EditICDRequest)
	if !ok {
		return response.InternalServerError("error from ur end, invalid context", nil)
	}

	icdFullCode := c.Params("fullCode")

	result, err := h.Repo.GetRecord(context.TODO(), icdFullCode)
	if err != nil {
		if entity.IsErrNotFound(err) {
			return response.
				NotFound(
					"the icd you are trying to edit does not exist", nil,
				)
		}

		return response.InternalServerError(err.Error(), nil)
	}

	result.AbbreviatedDescription = body.AbbreviatedDescription
	result.CategoryCode = body.CategoryCode
	result.CategoryTitle = body.CategoryTitle
	result.DiagnosisCode = body.DiagnosisCode
	result.FullDescription = body.FullDescription

	res, err := h.Repo.UpdateRecord(context.TODO(), icdFullCode, *result)
	if err != nil {
		return response.InternalServerError(err.Error(), nil)
	}

	resp := response.Ok("", res)
	return c.Status(resp.StatusCode).JSON(resp)
}

func (h *Handler) Delete(c *fiber.Ctx) error {

	icdFullCode := c.Params("fullCode")
	err := h.Repo.DeleteRecord(context.TODO(), icdFullCode)

	if err != nil {
		return response.InternalServerError(err.Error(), nil)
	}

	resp := response.Ok("", nil)
	return c.Status(resp.StatusCode).JSON(resp)
}

func (h *Handler) List(c *fiber.Ctx) error {
	cursor := c.Query("cursor", "0")
	limitPerPage := c.QueryInt("limit", 20)

	nextCursor, results, err := h.Repo.ListRecords(context.TODO(), limitPerPage, cursor)
	if err != nil {
		return response.InternalServerError(err.Error(), nil)
	}

	onFirstPage := cursor == "0"

	prevPage, err := h.Repo.HasPrevPage(context.TODO(), onFirstPage, cursor)
	if err != nil {
		return response.InternalServerError(err.Error(), nil)
	}

	resp := response.Ok("", fiber.Map{
		"prevPageCursor": prevPage,
		"nextPageCursor": nextCursor,
		"records":        results,
	})

	return c.Status(resp.StatusCode).JSON(resp)
}

func (h *Handler) Upload(c *fiber.Ctx) error {

	// Get first file from form field "document":
	file, err := c.FormFile("document")
	if err != nil {
		return response.InternalServerError(err.Error(), nil)
	}

	if !strings.HasSuffix(file.Filename, ".csv") {
		return response.BadRequest("only csv files are allowed", nil)
	}
	// Save file to root directory:
	err = c.SaveFile(file, fmt.Sprintf("%s/%s", h.Config.File.UploadDirectory, file.Filename))
	if err != nil {
		return response.BadRequest("error uploading file: "+err.Error(), nil)
	}

	// lets fire an event in a go-routine to notify email service
	go event.Fire(eventDTOs.TopicFileUploadComplete, event.M{"aa": "y@y.com"})

	resp := response.Ok("", nil)
	return c.Status(resp.StatusCode).JSON(resp)
}
