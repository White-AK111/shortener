package usecase

import (
	"github.com/White-AK111/shortener/internal/pkg/link"
)

type usecase struct {
	repo link.Repository
}

func New(repo link.Repository) link.Usecase {
	ret := usecase{
		repo: repo,
	}

	return ret
}
