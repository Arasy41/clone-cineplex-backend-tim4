package models

import (
	"time"

	"github.com/nsrvel/golang-example/pkg/db"
)

type Book struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Synopsis  *string    `json:"synopsis"`
	CoverUrl  *string    `json:"cover_url"`
	Content   string     `json:"content"`
	Author    string     `json:"author"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateBookRequest struct {
	Title    string `json:"title"`
	Synopsis string `json:"synopsis"`
	CoverUrl string `json:"cover_url"`
	Content  string `json:"content"`
	AuthorID int64  `json:"author_id"`
}

type GetAllBookRequest struct {
	AuthorID  *int64  `query:"filterAuthorId"`
	Search    *string `query:"search"`
	Page      int     `query:"page"`
	PageSize  int     `query:"pageSize"`
	OrderBy   string  `query:"orderBy"`
	OrderType string  `query:"orderType"`
}

type GetAllBookResponse struct {
	Books      []Book                 `json:"books"`
	Pagination *db.PaginationResponse `json:"pagination"`
}

type UpdateBookRequest struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Synopsis string `json:"synopsis"`
	CoverUrl string `json:"cover_url"`
	Content  string `json:"content"`
	AuthorID int64  `json:"author_id"`
}
