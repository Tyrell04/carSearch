package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type HelpHandler struct {
}

func NewHelpHandler() *HelpHandler {
	return &HelpHandler{}
}

func (h *HelpHandler) Route(api fiber.Router) {
	api.Get("/swagger/*", swagger.HandlerDefault) // default
}
