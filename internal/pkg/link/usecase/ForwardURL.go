package usecase

import (
	"context"
	"fmt"

	"github.com/White-AK111/shortener/internal/pkg/models"
)

func (u usecase) ForwardURL(ctx context.Context, link *models.Link) error {

	// do something
	if err := u.repo.FindURL(ctx, link); err != nil {
		return fmt.Errorf("failed to create item in repo: %w", err)
	}

	return nil
}
