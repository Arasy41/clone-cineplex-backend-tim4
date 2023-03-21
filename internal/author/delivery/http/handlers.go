package http

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nsrvel/golang-example/config"
	"github.com/nsrvel/golang-example/internal/author/usecase"
	"github.com/nsrvel/golang-example/internal/models"
	"github.com/nsrvel/golang-example/pkg/utils"
)

type AuthorHandler struct {
	uc  usecase.Usecase
	cfg *config.Config
}

func NewAuthorHandler(uc usecase.Usecase, cfg *config.Config) AuthorHandler {
	return AuthorHandler{
		uc:  uc,
		cfg: cfg,
	}
}

func (h AuthorHandler) CreateAuthor(c *fiber.Ctx) error {

	req := new(models.CreateAuthorRequest)
	if err := c.BodyParser(req); err != nil {
		return utils.WriteErrorResponse(c, http.StatusBadRequest, "Periksa kembali input anda !", err.Error())
	}

	//* Usecase
	err := h.uc.CreateAuthor(context.Background(), req)
	if err != nil {
		restErr := utils.ParseError(err)
		return utils.WriteErrorResponse(c, restErr.Status(), restErr.Message(), restErr.Error())
	}

	//* Create success response
	return utils.WriteSuccessResponse(c, http.StatusOK, "sukses", nil)
}

func (h AuthorHandler) GetAuthorByID(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteErrorResponse(c, http.StatusBadRequest, "pastikan id adalah sebuah integer", err.Error())
	}

	//* Usecase
	author, err := h.uc.GetAuthorByID(context.Background(), int64(id))
	if err != nil {
		restErr := utils.ParseError(err)
		return utils.WriteErrorResponse(c, restErr.Status(), restErr.Message(), restErr.Error())
	}

	//* Create success response
	return utils.WriteSuccessResponse(c, http.StatusOK, "sukses", author)
}

func (h AuthorHandler) GetAllAuthor(c *fiber.Ctx) error {

	//* Usecase
	authors, err := h.uc.GetAllAuthor(context.Background())
	if err != nil {
		restErr := utils.ParseError(err)
		return utils.WriteErrorResponse(c, restErr.Status(), restErr.Message(), restErr.Error())
	}

	//* Create success response
	return utils.WriteSuccessResponse(c, http.StatusOK, "sukses", authors)
}

func (h AuthorHandler) UpdateAuthor(c *fiber.Ctx) error {

	req := new(models.UpdateAuthorRequest)
	if err := c.BodyParser(req); err != nil {
		return utils.WriteErrorResponse(c, http.StatusBadRequest, "Periksa kembali input anda !", err.Error())
	}

	//* Usecase
	err := h.uc.UpdateAuthor(context.Background(), req)
	if err != nil {
		restErr := utils.ParseError(err)
		return utils.WriteErrorResponse(c, restErr.Status(), restErr.Message(), restErr.Error())
	}

	//* Create success response
	return utils.WriteSuccessResponse(c, http.StatusOK, "sukses", nil)
}

func (h AuthorHandler) DeleteAuthor(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteErrorResponse(c, http.StatusBadRequest, "pastikan id adalah sebuah integer", err.Error())
	}

	//* Usecase
	err = h.uc.DeleteAuthor(context.Background(), int64(id))
	if err != nil {
		restErr := utils.ParseError(err)
		return utils.WriteErrorResponse(c, restErr.Status(), restErr.Message(), restErr.Error())
	}

	//* Create success response
	return utils.WriteSuccessResponse(c, http.StatusOK, "sukses", nil)
}
