package handlers

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tomcyr/workos-go-authkit-example/application/common"
	"github.com/tomcyr/workos-go-authkit-example/conf"
	"github.com/tomcyr/workos-go-authkit-example/model/entity"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

type AuthHandler struct {
	creds conf.WorkOs
}

func NewAuthHandler(creds conf.WorkOs) *AuthHandler {
	return &AuthHandler{
		creds: creds,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	return c.Redirect().To(h.creds.AuthkitURL)
}

func (h *AuthHandler) Logout(c fiber.Ctx) error {
	user, err := common.GetUser(c)
	if err != nil {
		return err
	}
	usermanagement.SetAPIKey(
		h.creds.ApiKey,
	)

	err = usermanagement.RevokeSession(c.Context(), usermanagement.RevokeSessionOpts{SessionID: user.SID})
	if err != nil {
		return err
	}

	sess := session.FromContext(c)
	if sess != nil {
		err = sess.Destroy()
		if err != nil {
			return err
		}
	}

	return c.Redirect().To("/")
}

func (h *AuthHandler) Callback(c fiber.Ctx) error {
	usermanagement.SetAPIKey(
		h.creds.ApiKey,
	)

	resp, err := usermanagement.AuthenticateWithCode(c.Context(), usermanagement.AuthenticateWithCodeOpts{
		ClientID:  h.creds.ClientID,
		Code:      c.Query("code"),
		IPAddress: c.IP(),
		UserAgent: c.Get("User-Agent"),
	})

	if err != nil {
		return err
	}

	sess := session.FromContext(c)
	if sess == nil {
		return errors.New("session not found")
	}

	user := entity.NewUserFromAuth(resp.User.ID, resp.User.FirstName, resp.User.LastName, resp.User.Email, resp.User.EmailVerified)
	token, _, err := new(jwt.Parser).ParseUnverified(resp.AccessToken, jwt.MapClaims{})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		user.SID = claims["sid"].(string)
	}

	userBytes, _ := json.Marshal(user)
	sess.Set("user", string(userBytes))

	return c.Redirect().To("/dashboard")
}
