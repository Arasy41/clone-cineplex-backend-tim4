package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/nsrvel/golang-example/config"
	authorRepo "github.com/nsrvel/golang-example/internal/author/repository"
	"github.com/nsrvel/golang-example/internal/book/delivery/http"
	"github.com/nsrvel/golang-example/internal/book/repository"
	"github.com/nsrvel/golang-example/internal/book/usecase"
)

type Book struct {
	Repo    repository.Repository
	Usecase usecase.Usecase
	Handler http.BookHandler
}

func NewBook(cfg *config.Config, authorRepo authorRepo.Repository, sqlDB *sqlx.DB) Book {
	repo := repository.NewBookRepo(sqlDB)
	uc := usecase.NewBookUsecase(repo, authorRepo, cfg)
	handler := http.NewBookHandler(uc, cfg)
	return Book{Repo: repo, Usecase: uc, Handler: handler}
}

func NewRoutes(r fiber.Router, handler http.BookHandler) {
	http.MapBookRoutes(r, handler)
}
