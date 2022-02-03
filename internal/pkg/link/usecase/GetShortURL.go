package usecase

import (
	"context"

	"github.com/White-AK111/shortener/internal/pkg/models"
)

func (u usecase) GetShortURL(ctx context.Context, link *models.Link) error {
	// do something
	return nil
}
