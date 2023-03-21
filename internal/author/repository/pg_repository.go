package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nsrvel/golang-example/internal/models"
	"github.com/nsrvel/golang-example/pkg/db"
)

type Repository interface {
	CreateAuthor(ctx context.Context, params ...interface{}) (int64, error)
	GetAuthorByID(ctx context.Context, id int64) (*models.Author, error)
	GetAuthorByName(ctx context.Context, name string) (*models.Author, error)
	GetAllAuthor(ctx context.Context) (*[]models.Author, error)
	UpdateAuthor(ctx context.Context, params ...interface{}) error
	DeleteAuthor(ctx context.Context, id int64) error
}

type AuthorRepo struct {
	sqlDB *sqlx.DB
}

func NewAuthorRepo(sqlDB *sqlx.DB) AuthorRepo {
	return AuthorRepo{
		sqlDB: sqlDB,
	}
}

func (r AuthorRepo) CreateAuthor(ctx context.Context, params ...interface{}) (int64, error) {
	var response int64
	response, err := db.QueryRow(ctx, r.sqlDB, response, qCreateAuthor, params...)
	return response, err
}

func (r AuthorRepo) GetAuthorByID(ctx context.Context, id int64) (*models.Author, error) {
	var response models.Author
	response, err := db.QueryRow(ctx, r.sqlDB, response, qGetAuthorByID, id)
	return &response, err
}

func (r AuthorRepo) GetAuthorByName(ctx context.Context, name string) (*models.Author, error) {
	var response models.Author
	response, err := db.QueryRow(ctx, r.sqlDB, response, qGetAuthorByName, name)
	return &response, err
}

func (r AuthorRepo) GetAllAuthor(ctx context.Context) (*[]models.Author, error) {
	var response []models.Author
	response, err := db.Query(ctx, r.sqlDB, response, qGetAllAuthor)
	return &response, err
}

func (r AuthorRepo) UpdateAuthor(ctx context.Context, params ...interface{}) error {
	err := db.Exec(ctx, r.sqlDB, qUpdateAuthor, params...)
	return err
}

func (r AuthorRepo) DeleteAuthor(ctx context.Context, id int64) error {
	err := db.Exec(ctx, r.sqlDB, qDeleteAuthor, id)
	return err
}
