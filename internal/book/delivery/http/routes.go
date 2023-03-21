package http

import "github.com/gofiber/fiber/v2"

func MapBookRoutes(r fiber.Router, handler BookHandler) {

	books := r.Group("/books")

	books.Post("/", handler.CreateBook)
	books.Get("/:id", handler.GetBookByID)
	books.Get("/", handler.GetAllBook)
	books.Put("/", handler.UpdateBook)
	books.Delete("/:id", handler.DeleteBook)
}
