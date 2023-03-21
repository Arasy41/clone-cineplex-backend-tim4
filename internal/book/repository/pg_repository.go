package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nsrvel/golang-example/internal/models"
	"github.com/nsrvel/golang-example/pkg/db"
)

type Repository interface {
	CreateBook(ctx context.Context, params ...interface{}) (int64, error)
	GetBookByID(ctx context.Context, id int64) (*models.Book, error)
	GetAllBook(ctx context.Context, filterAuthorID *int64, searchValue *string, paging *db.PaginationRequest) (*[]models.Book, *db.PaginationResponse, error)
	UpdateBook(ctx context.Context, params ...interface{}) error
	DeleteBook(ctx context.Context, id int64) error
}

type BookRepo struct {
	sqlDB *sqlx.DB
}

func NewBookRepo(sqlDB *sqlx.DB) BookRepo {
	return BookRepo{
		sqlDB: sqlDB,
	}
}

func (r BookRepo) CreateBook(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	response, err := db.QueryRow(ctx, r.sqlDB, response, qCreateBook, params...)
	return response, err
}

func (r BookRepo) GetBookByID(ctx context.Context, id int64) (*models.Book, error) {
	var response models.Book
	response, err := db.QueryRow(ctx, r.sqlDB, response, qGetBookByID, id)
	return &response, err
}

//* Yang ini lumayan advance, bisa kalian liat & ikuti dari cara penggunaannya saja
func (r BookRepo) GetAllBook(ctx context.Context, filterAuthorID *int64, searchValue *string, paging *db.PaginationRequest) (*[]models.Book, *db.PaginationResponse, error) {

	var response []models.Book
	var total int
	var err error

	//* Setingan filter
	filter := db.FilterData{

		//* Untuk filter
		Filter: &[]db.Filter{
			{
				Key:      "author_id",
				Value:    filterAuthorID,
				Operator: "=",
			},
		},

		//* Untuk Search
		Search: &[]db.Search{
			{
				Value:    searchValue,
				Operator: "ILIKE",
				Target: []string{
					"title",
				},
			},
		},
	}
	qFilter := filter.FilterQueryBuilder(true)

	//* Setingan Pagination
	page := db.PageData{

		//* Data pagination
		Paging: paging,

		//* Untuk Setingan Defaultnya juga
		Default: db.Default{

			//* 10 data akan tampil
			Size: 10,

			//* Data akan diurutkan berdasarkan column created_at
			OrderBy: "created_at",

			//* Data akan diurutkan dari nilai tertinggi ke nilai terendah
			OrderType: "DESC",
		},
	}
	qPage := page.PaginationQueryBuilder(false)

	//* Mencari total data sebenarnya
	total, err = db.QueryRow(ctx, r.sqlDB, total, qCountBooks+qFilter)
	if err != nil {
		return nil, nil, err
	}

	response, err = db.Query(ctx, r.sqlDB, response, qGetAllBook+qFilter+qPage)

	return &response, page.GetPaginationResponse(total), err
}

func (r BookRepo) UpdateBook(ctx context.Context, params ...interface{}) error {
	return db.Exec(ctx, r.sqlDB, qUpdateBook, params...)
}

func (r BookRepo) DeleteBook(ctx context.Context, id int64) error {
	return db.Exec(ctx, r.sqlDB, qDeleteBook, id)
}
