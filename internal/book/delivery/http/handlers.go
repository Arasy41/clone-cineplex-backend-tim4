package http

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nsrvel/golang-example/config"
	"github.com/nsrvel/golang-example/internal/book/usecase"
	"github.com/nsrvel/golang-example/internal/models"
	"github.com/nsrvel/golang-example/pkg/utils"
)

type BookHandler struct {
	uc  usecase.Usecase
	cfg *config.Config
}

func NewBookHandler(uc usecase.Usecase, cfg *config.Config) BookHandler {
	return BookHandler{
		uc:  uc,
		cfg: cfg,
	}
}

func (h BookHandler) CreateBook(c *fiber.Ctx) error {

	req := new(models.CreateBookRequest)
	if err := c.BodyParser(req); err != nil {
		return utils.WriteErrorResponse(c, http.StatusBadRequest, "Periksa kembali input anda !", err.Error())
	}

	//* Usecase
	err := h.uc.CreateBook(context.Background(), req)
	if err != nil {
		restErr := utils.ParseError(err)
		return utils.WriteErrorResponse(c, restErr.Status(), restErr.Message(), restErr.Error())
	}

	//* Create success response
	return utils.WriteSuccessResponse(c, http.StatusOK, "sukses", nil)
}

func (h BookHandler) GetBookByID(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteErrorResponse(c, http.StatusBadRequest, "pastikan id adalah sebuah integer", err.Error())
	}

	//* Usecase
	author, err := h.uc.GetBookByID(context.Background(), int64(id))
	if err != nil {
		restErr := utils.ParseError(err)
		return utils.WriteErrorResponse(c, restErr.Status(), restErr.Message(), restErr.Error())
	}

	//* Create success response
	return utils.WriteSuccessResponse(c, http.StatusOK, "sukses", author)
}

func (h BookHandler) GetAllBook(c *fiber.Ctx) error {

	req := new(models.GetAllBookRequest)
	if err := c.QueryParser(req); err != nil {
		return utils.WriteErrorResponse(c, http.StatusBadRequest, "Periksa kembali input anda !", err.Error())
	}

	//* Usecase
	authors, err := h.uc.GetAllBook(context.Background(), req)
	if err != nil {
		restErr := utils.ParseError(err)
		return utils.WriteErrorResponse(c, restErr.Status(), restErr.Message(), restErr.Error())
	}

	//* Create success response
	return utils.WriteSuccessResponse(c, http.StatusOK, "sukses", authors)
}

func (h BookHandler) UpdateBook(c *fiber.Ctx) error {

	req := new(models.UpdateBookRequest)
	if err := c.BodyParser(req); err != nil {
		return utils.WriteErrorResponse(c, http.StatusBadRequest, "Periksa kembali input anda !", err.Error())
	}

	//* Usecase
	err := h.uc.UpdateBook(context.Background(), req)
	if err != nil {
		restErr := utils.ParseError(err)
		return utils.WriteErrorResponse(c, restErr.Status(), restErr.Message(), restErr.Error())
	}

	//* Create success response
	return utils.WriteSuccessResponse(c, http.StatusOK, "sukses", nil)
}

func (h BookHandler) DeleteBook(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteErrorResponse(c, http.StatusBadRequest, "pastikan id adalah sebuah integer", err.Error())
	}

	//* Usecase
	err = h.uc.DeleteBook(context.Background(), int64(id))
	if err != nil {
		restErr := utils.ParseError(err)
		return utils.WriteErrorResponse(c, restErr.Status(), restErr.Message(), restErr.Error())
	}

	//* Create success response
	return utils.WriteSuccessResponse(c, http.StatusOK, "sukses", nil)
}
