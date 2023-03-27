package route

import (
	"github.com/Artzy1401/golang-restapi-gorm/config"
	"github.com/Artzy1401/golang-restapi-gorm/controller"
	"github.com/Artzy1401/golang-restapi-gorm/middleware"
	"github.com/gofiber/fiber/v2"
)


func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", controller.LoginController)

	r.Get("/film", controller.FilmControllerGetAll)
	r.Get("/film/:id", controller.FilmControllerGetById)
	r.Post("/film", controller.FilmControllerCreate)
	r.Put("/film/:id", controller.FilmControllerUpdate)
	r.Delete("/film/:id", controller.FilmControllerDelete)

	r.Get("/user", middleware.Auth, controller.UserControllerGetAll)
	r.Get("/user/:id", controller.UserControllerGetById)
	r.Post("/user", controller.UserControllerCreate)
	r.Put("/user/:id", controller.UserControllerUpdate)
	// r.Put("/user/:id/update-email", controller.UserControllerUpdateEmail)
	r.Delete("/user/:id", controller.UserControllerDelete)
}