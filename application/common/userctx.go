package common

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/tomcyr/workos-go-authkit-example/model/entity"
)

func GetUser(c fiber.Ctx) (*entity.User, error) {
	sess := session.FromContext(c)
	if sess == nil {
		return nil, errors.New("session not found")
	}

	userStr := sess.Get("user")
	if userStr == nil {
		return nil, errors.New("user not found in session")
	}

	var user entity.User
	err := json.Unmarshal([]byte(userStr.(string)), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
