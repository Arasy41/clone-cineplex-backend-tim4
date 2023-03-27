package controller

import (
	"log"

	"github.com/Artzy1401/golang-restapi-gorm/database"
	"github.com/Artzy1401/golang-restapi-gorm/model/entity"
	"github.com/Artzy1401/golang-restapi-gorm/model/request"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func FilmControllerGetAll(ctx *fiber.Ctx) error {
	var film []entity.Film
	result := database.DB.Find(&film)
	if result.Error != nil {
		log.Println(result.Error)
	}

	// err := database.DB.Find(&film).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return ctx.JSON(film)
}

func FilmControllerCreate(ctx *fiber.Ctx) error {
	film := new(request.FilmCreateRequest)
	if err := ctx.BodyParser(film); err != nil {
		return err
	}

	// VALIDASI REQUEST
 	validate := validator.New()
	errValidate := validate.Struct(film)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "failed to validate",
			"error" : errValidate.Error(),
		})
	}


	newFilm := entity.Film{
		Name:		film.Name,
		JenisFilm: 	film.JenisFilm,
		Produser: 	film.Produser,
		Sutradara: 	film.Sutradara,
		Penulis:	film.Penulis,
		Produksi:	film.Produksi,
		Casts:		film.Casts,
		Sinopsis:	film.Sinopsis,
	}

	errCreateUser := database.DB.Create(&newFilm).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to create user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": newFilm,
	})
}


func FilmControllerGetById(ctx *fiber.Ctx) error{
	filmId := ctx.Params("id")

	var film entity.Film
	err := database.DB.First(&film, "id = ?", filmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": film,
	})
}

func FilmControllerUpdate (ctx *fiber.Ctx) error {
	filmRequest := new(request.FilmUpdateRequest)
	if err := ctx.BodyParser(filmRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var film entity.Film

	filmId := ctx.Params("id")
	// CHECK AVALAIBLE FILM
	err := database.DB.First(&film, "id = ?", filmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	// UPDATE USER DATA
	if filmRequest.Name != "" {
		film.Name = filmRequest.Name
	}
	film.JenisFilm = filmRequest.JenisFilm
	film.Produksi = filmRequest.Produksi
	film.Sutradara = filmRequest.Sutradara
	film.Penulis = filmRequest.Penulis
	film.Produksi =	filmRequest.Produksi
	film.Casts = filmRequest.Casts
	film.Sinopsis =	filmRequest.Sinopsis

	errUpdate := database.DB.Save(&film).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": film,
	})
}

func FilmControllerDelete(ctx *fiber.Ctx) error {
	filmId := ctx.Params("id")
	var film entity.Film

	// CHECK AVAILABLE FILM
	err := database.DB.Debug().First(&film, "id=?" ,&filmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Film Not Found",
		})
	}

	errDelete := database.DB.Debug().Delete(&film).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "film deleted",
	})
}