package author

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/nsrvel/golang-example/config"
	"github.com/nsrvel/golang-example/internal/author/delivery/http"
	"github.com/nsrvel/golang-example/internal/author/repository"
	"github.com/nsrvel/golang-example/internal/author/usecase"
)

type Author struct {
	Repo    repository.Repository
	Usecase usecase.Usecase
	Handler http.AuthorHandler
}

func NewAuthor(cfg *config.Config, sqlDB *sqlx.DB) Author {
	repo := repository.NewAuthorRepo(sqlDB)
	uc := usecase.NewAuthorUsecase(repo, cfg)
	handler := http.NewAuthorHandler(uc, cfg)
	return Author{Repo: repo, Usecase: uc, Handler: handler}
}

func NewRoutes(r fiber.Router, handler http.AuthorHandler) {
	http.MapAuthorRoutes(r, handler)
}
