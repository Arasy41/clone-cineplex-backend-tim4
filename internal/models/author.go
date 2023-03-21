package models

import "time"

type Author struct {
	ID        int64      `json:"id"`
	FullName  string     `json:"full_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateAuthorRequest struct {
	FullName string `json:"full_name"`
}

type UpdateAuthorRequest struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
}
