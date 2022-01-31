package postgres

import (
	"context"

	"github.com/White-AK111/shortener/internal/pkg/models"
)

func (r repository) IncrementStat(ctx context.Context, link *models.Link) (int, error) {
	// do something
	return 0, nil
}
