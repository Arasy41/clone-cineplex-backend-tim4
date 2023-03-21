package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/nsrvel/golang-example/config"
	"github.com/nsrvel/golang-example/internal/author"
	"github.com/nsrvel/golang-example/internal/book"
	"github.com/nsrvel/golang-example/pkg/utils"
)

func RunFiber(cfg *config.Config, sqlDB *sqlx.DB) {

	app := fiber.New(fiber.Config{
		AppName:      cfg.Server.Name,
		ServerHeader: "Go Fiber",
	})

	//* panggil domain
	Author := author.NewAuthor(cfg, sqlDB)
	Book := book.NewBook(cfg, Author.Repo, sqlDB)

	//* root
	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("Hello World") })

	//* v1
	v1 := app.Group("/api/v1")

	//* Routes
	author.NewRoutes(v1, Author.Handler)
	book.NewRoutes(v1, Book.Handler)

	//* Not found
	app.All("*", func(c *fiber.Ctx) error {
		return utils.WriteErrorResponse(c, http.StatusNotFound, "Ups! Kami tidak dapat menemukan apa yang Anda cari.", "")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)))
}
