package middleware

import "grob/internal/repository"

type Middleware struct{}

func New(repo *repository.Repository) *Middleware {
	return &Middleware{}
}
