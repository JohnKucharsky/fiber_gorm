package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"os/signal"
)

func New() *fiber.App {
	f := fiber.New()
	f.Use(logger.New())
	f.Use(
		cors.New(
			cors.Config{
				AllowOrigins: "*",
				AllowHeaders: "Origin, Content-Type, Accept, Authorization",
				AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
			},
		),
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = f.Shutdown()
	}()

	return f
}
