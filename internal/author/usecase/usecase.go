package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nsrvel/golang-example/config"
	"github.com/nsrvel/golang-example/internal/author/repository"
	"github.com/nsrvel/golang-example/internal/models"
	"github.com/nsrvel/golang-example/pkg/utils"
)

type Usecase interface {
	CreateAuthor(ctx context.Context, req *models.CreateAuthorRequest) error
	GetAuthorByID(ctx context.Context, id int64) (*models.Author, error)
	GetAllAuthor(ctx context.Context) (*[]models.Author, error)
	UpdateAuthor(ctx context.Context, req *models.UpdateAuthorRequest) error
	DeleteAuthor(ctx context.Context, id int64) error
}

type AuthorUsecase struct {
	repo repository.Repository
	cfg  *config.Config
}

func NewAuthorUsecase(repo repository.Repository, cfg *config.Config) AuthorUsecase {
	return AuthorUsecase{
		repo: repo,
		cfg:  cfg,
	}
}

func (u AuthorUsecase) CreateAuthor(ctx context.Context, req *models.CreateAuthorRequest) error {

	var err error

	//* Validasi data request
	if req.FullName == "" {
		return utils.ErrorWrapper(http.StatusBadRequest, "data full_name tidak boleh kosong", nil)
	}

	//* Cek author duplicate atau tidak
	author, err := u.repo.GetAuthorByName(ctx, req.FullName)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data author", err)
	}
	if author != nil {
		if author.FullName == req.FullName {
			return utils.ErrorWrapper(http.StatusBadRequest, "author sudah ada", nil)
		}
	}

	params := make([]interface{}, 0)
	params = append(params, req.FullName)

	authorID, err := u.repo.CreateAuthor(ctx, params...)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal insert data author", err)
	}
	if authorID == 0 {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal insert data author", nil)
	}

	return nil
}

func (u AuthorUsecase) GetAuthorByID(ctx context.Context, id int64) (*models.Author, error) {

	var err error

	//* Validasi data request
	if id == 0 {
		return nil, utils.ErrorWrapper(http.StatusBadRequest, "data id tidak boleh kosong", nil)
	}

	author, err := u.repo.GetAuthorByID(ctx, id)
	if err != nil {
		return nil, utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data author", err)
	}
	if author != nil {
		if author.ID == 0 {
			return nil, utils.ErrorWrapper(http.StatusNotFound, "author tidak ditemukan", nil)
		}
	}

	return author, nil
}

func (u AuthorUsecase) GetAllAuthor(ctx context.Context) (*[]models.Author, error) {

	var err error

	author, err := u.repo.GetAllAuthor(ctx)
	if err != nil {
		return nil, utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data author", err)
	}
	if len(*author) == 0 {
		fmt.Println("data author tidak ditemukan")
		return nil, nil
	}

	return author, nil
}

func (u AuthorUsecase) UpdateAuthor(ctx context.Context, req *models.UpdateAuthorRequest) error {

	var err error

	//* Validasi data request
	if req.ID == 0 {
		return utils.ErrorWrapper(http.StatusBadRequest, "data id tidak boleh kosong", nil)
	}
	if req.FullName == "" {
		return utils.ErrorWrapper(http.StatusBadRequest, "data full_name tidak boleh kosong", nil)
	}

	//* cek author ada atau tidak
	author, err := u.repo.GetAuthorByID(ctx, req.ID)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data author", err)
	}
	if author != nil {
		if author.ID == 0 {
			return utils.ErrorWrapper(http.StatusNotFound, "author tidak ditemukan", nil)
		}
	}

	//* Cek author duplicate atau tidak
	dupe, err := u.repo.GetAuthorByName(ctx, req.FullName)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data author", err)
	}
	if dupe != nil {
		if req.FullName != author.FullName {
			if dupe.FullName == req.FullName {
				return utils.ErrorWrapper(http.StatusBadRequest, "author sudah ada", nil)
			}
		}
	}

	params := make([]interface{}, 0)
	params = append(params, req.FullName, req.ID)

	err = u.repo.UpdateAuthor(ctx, params...)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal update data author", err)
	}

	return nil
}

func (u AuthorUsecase) DeleteAuthor(ctx context.Context, id int64) error {

	var err error

	//* Validasi data request
	if id == 0 {
		return utils.ErrorWrapper(http.StatusBadRequest, "data id tidak boleh kosong", nil)
	}

	//* cek author ada atau tidak
	author, err := u.repo.GetAuthorByID(ctx, id)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data author", err)
	}
	if author != nil {
		if author.ID == 0 {
			return utils.ErrorWrapper(http.StatusNotFound, "author tidak ditemukan", nil)
		}
	}

	err = u.repo.DeleteAuthor(ctx, id)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal delete data author", err)
	}

	return nil
}
