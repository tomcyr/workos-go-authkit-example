package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/tomcyr/workos-go-authkit-example/application/handlers"
	"github.com/tomcyr/workos-go-authkit-example/application/middleware"
	"github.com/tomcyr/workos-go-authkit-example/conf"
)

var configFile = flag.String(
	"config_file",
	"conf/config.yaml",
	"Path to the YAML config",
)

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT,
	)
	defer cancel()

	cfg, err := conf.ParseConfig(*configFile)
	if err != nil {
		panic(err)
	}

	sessionMiddleware, sessionStore := session.NewWithStore(
		session.Config{
			KeyLookup: "cookie:workos-example",
		})

	engine := html.New("./views", ".html")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)
	app.Use(requestid.New(), recover.New())
	app.Use(sessionMiddleware)
	app.Use(csrf.New(csrf.Config{
		Session: sessionStore,
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	homepageHandler := handlers.NewHomePageHandler()
	authHandler := handlers.NewAuthHandler(cfg.WorkOs)
	dashboardHandler := handlers.NewDashboardHandler()

	defaultG := app.Group("")
	defaultG.Get("/", homepageHandler.Index)
	defaultG.Get("/auth/login", authHandler.Login)
	defaultG.Get("/auth/callback", authHandler.Callback)
	defaultG.Get("/auth/logout", authHandler.Logout, middleware.AuthMiddleware())

	dashboardG := app.Group("/dashboard", middleware.AuthMiddleware())
	dashboardG.Get("/", dashboardHandler.Index)

	log.Fatal(app.Listen(cfg.HTTP.Address, fiber.ListenConfig{
		GracefulContext: ctx,
	}))
}
