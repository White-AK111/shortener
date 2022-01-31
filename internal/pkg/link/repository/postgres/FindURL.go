package postgres

import (
	"context"

	"github.com/White-AK111/shortener/internal/pkg/models"
)

func (r repository) FindURL(ctx context.Context, link *models.Link) error {
	// do something
	return nil
}
