package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/HOA-Share/hoashare-go-common/http_server"
	"github.com/HOA-Share/hoashare-go-common/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/pranavpatil6/config"
	routes "github.com/pranavpatil6/public"
)

func main() {
	config.Init()
	logger.GetLogger().Info("Starting server")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	app := http_server.NewRouter(http_server.RouterConfig{
		AppName:           "Uber-Anylitics	",
		StrictRouting:     true,
		IsRecoveryEnabled: true,
		IsLoggerEnabled:   true,
		Port:              config.GetConfig().ServerConfig.Port,
		GlobalCors:        []string{},
	})

	app.Get("/hi", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "Doneeeee",
		})
	})
	app.Get("/debug/routes", func(c *fiber.Ctx) error {
		data := app.GetRoutes()
		return c.JSON(data)
		
	})

	routes.MountRoutes(app)

	app.StartAsync()
	logger.GetLogger().Info("server started")
	<-sig
}
