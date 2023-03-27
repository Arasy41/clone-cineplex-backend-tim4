package request

import (
	
)

type FilmCreateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	JenisFilm 	string `json:"jenis_film" validate:"required"`
	Produser	string `json:"produser" validate:"required"`
	Sutradara	string `json:"sutradara"`
	Penulis		string `json:"penulis"`
	Produksi	string `json:"produksi"`
	Casts		string `json:"casts"`
	Sinopsis	string `json:"sinopsis"`
}

type FilmUpdateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	JenisFilm 	string `json:"jenis_film"`
	Produser	string `json:"produser"`
	Sutradara	string `json:"sutradara"`
	Penulis		string `json:"penulis"`
	Produksi	string `json:"produksi"`
	Casts		string `json:"casts"`
	Sinopsis	string `json:"sinopsis"`
}

type UserCreateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	Email 		string 	`json:"email" validate:"required,email"`
	Password 	string 	`json:"password" validate:"required,min=6"`
}

type UserUpdateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	Email 		string 	`json:"email" validate:"required"`
	Password 	string 	`json:"password" validate:"required"`
}

type UserEmailRequest struct {
	Email 		string 	`json:"email" validate:"required"`
}