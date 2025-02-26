package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tomcyr/workos-go-authkit-example/application/common"
)

type DashboardHandler struct {
}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{}
}

func (h *DashboardHandler) Index(c fiber.Ctx) error {
	username := ""
	user, _ := common.GetUser(c)
	if user != nil {
		username = user.Email
	}

	return c.Render("dashboard/index", fiber.Map{
		"Authenticated": true,
		"Username":      username,
	}, "layouts/main")
}
