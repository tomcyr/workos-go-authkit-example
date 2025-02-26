package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		sess := session.FromContext(c)
		if sess == nil {
			return c.Redirect().To("/auth/login")
		}

		user := sess.Get("user")
		if user == nil {
			return c.Redirect().To("/auth/login")
		}

		return c.Next()
	}
}
