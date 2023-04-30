package middleware

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/otyang/icd-10/pkg/logger"
	"github.com/otyang/icd-10/pkg/response"
)

func HttpLog(log logger.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			start      = time.Now()
			err        = c.Next()
			duration   = time.Since(start).Milliseconds()
			statusCode = getStatusCode(err, c)
		)

		httpLogParams := logger.HTTPLogParams{
			Error:           err,
			StatusCode:      statusCode,
			Duration:        duration,
			Method:          c.Method(),
			RequestID:       getRequestID(c),
			ErrorCallerInfo: getErrRuntimeCallerInfo(err),
			UserAgent:       c.GetReqHeaders()["User-Agent"],
			Link:            string(c.Request().URI().FullURI()),
		}

		if err != nil {
			logger.HTTPLog(log, httpLogParams, "error-returned-from-handler")
			// since the error was returned from handler lets format json response
			return c.Status(statusCode).JSON(err)
		}

		if statusCode >= http.StatusInternalServerError {
			logger.HTTPLog(log, httpLogParams, "error-triggered-on-http-status-500-and-above")
		}

		return nil
	}
}

func getRequestID(c *fiber.Ctx) string {
	reqHeaderID := c.GetReqHeaders()["X-Request-Id"]
	if reqHeaderID == "" {
		return "unknown"
	}
	return reqHeaderID
}

func getErrRuntimeCallerInfo(err error) string {
	if err != nil {
		if rsp, ok := err.(*response.Response); ok {
			return rsp.RuntimeCallerInfo
		}
		return "unknown"
	}
	return ""
}

func getStatusCode(err error, c *fiber.Ctx) int {
	if err != nil {
		if rsp, ok := err.(*response.Response); ok {
			return rsp.StatusCode
		}
	}
	return c.Response().StatusCode()
}
