package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valikhanyeltay/cms_api/internal/models"
)

func (h *Handler) CreateContentType(c *fiber.Ctx) (err error) {
	var payload models.ContentType

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	err = h.service.CreateContentType(&payload)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(models.SuccessResponse{
		Status:  true,
		Message: "SUCCESS",
	})
}
