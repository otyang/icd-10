package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leebenson/conform"
	"github.com/otyang/icd-10/pkg/response"
	"github.com/otyang/icd-10/pkg/validators"
)

func ValidateBody[BodyType any](c *fiber.Ctx) error {
	body := new(BodyType)
	if err := c.BodyParser(body); err != nil {
		rsp := response.BadRequest(err.Error(), nil)
		return c.Status(rsp.StatusCode).JSON(rsp)
	}

	conform.Strings(body) // so we can trim the spaces and apply changes

	err := validators.NewCheckerGoKit(body)
	msg := validators.TranslatorGoKit(err)

	if err != nil {
		rsp := response.BadRequest(msg, nil)
		return c.Status(rsp.StatusCode).JSON(rsp)
	}

	c.Locals("jsonBody", body)
	return c.Next()
}

// ValidatedDataFromContext(c).(*type)
func ValidatedDataFromContext(ctx *fiber.Ctx) (val any) {
	return ctx.Locals("jsonBody")
}
