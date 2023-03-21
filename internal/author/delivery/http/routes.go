package http

import "github.com/gofiber/fiber/v2"

func MapAuthorRoutes(r fiber.Router, handler AuthorHandler) {

	authors := r.Group("/authors")

	authors.Post("/", handler.CreateAuthor)
	authors.Get("/:id", handler.GetAuthorByID)
	authors.Get("/", handler.GetAllAuthor)
	authors.Put("/", handler.UpdateAuthor)
	authors.Delete("/:id", handler.DeleteAuthor)
}
