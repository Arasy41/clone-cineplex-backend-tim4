package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nsrvel/golang-example/config"
	authorRepo "github.com/nsrvel/golang-example/internal/author/repository"
	"github.com/nsrvel/golang-example/internal/book/repository"
	"github.com/nsrvel/golang-example/internal/models"
	"github.com/nsrvel/golang-example/pkg/db"
	"github.com/nsrvel/golang-example/pkg/utils"
)

type Usecase interface {
	CreateBook(ctx context.Context, req *models.CreateBookRequest) error
	GetBookByID(ctx context.Context, id int64) (*models.Book, error)
	GetAllBook(ctx context.Context, req *models.GetAllBookRequest) (*models.GetAllBookResponse, error)
	UpdateBook(ctx context.Context, req *models.UpdateBookRequest) error
	DeleteBook(ctx context.Context, id int64) error
}

type BookUsecase struct {
	repo       repository.Repository
	authorRepo authorRepo.Repository
	cfg        *config.Config
}

func NewBookUsecase(repo repository.Repository, authorRepo authorRepo.Repository, cfg *config.Config) BookUsecase {
	return BookUsecase{
		repo:       repo,
		authorRepo: authorRepo,
		cfg:        cfg,
	}
}

func (u BookUsecase) CreateBook(ctx context.Context, req *models.CreateBookRequest) error {

	var err error

	//* Validasi data request
	if req.Title == "" {
		return utils.ErrorWrapper(http.StatusBadRequest, "data title tidak boleh kosong", nil)
	}
	if req.Content == "" {
		return utils.ErrorWrapper(http.StatusBadRequest, "data content tidak boleh kosong", nil)
	}
	if req.AuthorID == 0 {
		return utils.ErrorWrapper(http.StatusBadRequest, "data author_id tidak boleh kosong", nil)
	}

	//* Cek author id tersedia atau tidak
	author, err := u.authorRepo.GetAuthorByID(ctx, req.AuthorID)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data author", err)
	}
	if author != nil {
		if author.ID == 0 {
			return utils.ErrorWrapper(http.StatusNotFound, "author tidak ditemukan", nil)
		}
	}

	params := make([]interface{}, 0)
	params = append(params, req.Title, req.Synopsis, req.CoverUrl, req.Content, req.AuthorID)

	bookID, err := u.repo.CreateBook(ctx, params...)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal insert data book", err)
	}
	if bookID == 0 {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal insert data book", nil)
	}

	return nil
}

func (u BookUsecase) GetBookByID(ctx context.Context, id int64) (*models.Book, error) {

	var err error

	//* Validasi data request
	if id == 0 {
		return nil, utils.ErrorWrapper(http.StatusBadRequest, "data id tidak boleh kosong", nil)
	}

	//* Cek author id tersedia atau tidak
	book, err := u.repo.GetBookByID(ctx, id)
	if err != nil {
		return nil, utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data book", err)
	}
	if book != nil {
		if book.ID == 0 {
			return nil, utils.ErrorWrapper(http.StatusNotFound, "book tidak ditemukan", nil)
		}
	}

	return book, nil
}

func (u BookUsecase) GetAllBook(ctx context.Context, req *models.GetAllBookRequest) (*models.GetAllBookResponse, error) {

	var err error

	//* Cek author id tersedia atau tidak
	books, paging, err := u.repo.GetAllBook(ctx, req.AuthorID, req.Search, &db.PaginationRequest{
		Page:      req.Page,
		Size:      req.PageSize,
		OrderBy:   req.OrderBy,
		OrderType: req.OrderType,
	})
	if err != nil {
		return nil, utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data book", err)
	}
	if len(*books) == 0 {
		fmt.Println("data book tidak ditemukan")
	}

	return &models.GetAllBookResponse{
		Books:      *books,
		Pagination: paging,
	}, nil
}

func (u BookUsecase) UpdateBook(ctx context.Context, req *models.UpdateBookRequest) error {

	var err error

	//* Validasi data request
	if req.ID == 0 {
		return utils.ErrorWrapper(http.StatusBadRequest, "data id tidak boleh kosong", nil)
	}
	if req.Title == "" {
		return utils.ErrorWrapper(http.StatusBadRequest, "data title tidak boleh kosong", nil)
	}
	if req.Content == "" {
		return utils.ErrorWrapper(http.StatusBadRequest, "data content tidak boleh kosong", nil)
	}
	if req.AuthorID == 0 {
		return utils.ErrorWrapper(http.StatusBadRequest, "data author_id tidak boleh kosong", nil)
	}

	//* cek book ada atau tidak
	author, err := u.repo.GetBookByID(ctx, req.ID)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data book", err)
	}
	if author != nil {
		if author.ID == 0 {
			return utils.ErrorWrapper(http.StatusNotFound, "book tidak ditemukan", nil)
		}
	}

	params := make([]interface{}, 0)
	params = append(params, req.Title, req.Synopsis, req.CoverUrl, req.Synopsis, req.AuthorID, req.ID)

	err = u.repo.UpdateBook(ctx, params...)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal update data book", err)
	}

	return nil
}

func (u BookUsecase) DeleteBook(ctx context.Context, id int64) error {

	var err error

	//* Validasi data request
	if id == 0 {
		return utils.ErrorWrapper(http.StatusBadRequest, "data id tidak boleh kosong", nil)
	}

	//* cek author ada atau tidak
	author, err := u.repo.GetBookByID(ctx, id)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal get data book", err)
	}
	if author != nil {
		if author.ID == 0 {
			return utils.ErrorWrapper(http.StatusNotFound, "book tidak ditemukan", nil)
		}
	}

	err = u.repo.DeleteBook(ctx, id)
	if err != nil {
		return utils.ErrorWrapper(http.StatusInternalServerError, "gagal delete data book", err)
	}

	return nil
}
