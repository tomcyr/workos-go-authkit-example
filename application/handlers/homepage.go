package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tomcyr/workos-go-authkit-example/application/common"
)

type HomePageHandler struct {
}

func NewHomePageHandler() *HomePageHandler {
	return &HomePageHandler{}
}

func (h *HomePageHandler) Index(c fiber.Ctx) error {
	authenticated := false
	user, _ := common.GetUser(c)
	if user != nil {
		authenticated = true
	}

	return c.Render("homepage/index", fiber.Map{
		"Authenticated": authenticated,
	}, "layouts/main")
}
