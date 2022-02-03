package link

import (
	"context"
	"errors"

	"github.com/White-AK111/shortener/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

var (
	ErrLinkNotFound = errors.New("link not found")
)

type GinDelivery interface {
	ForwardURL(ctx *gin.Context)
	GetShortURL(ctx *gin.Context)
	GetStat(ctx *gin.Context)
}

type Usecase interface {
	ForwardURL(ctx context.Context, link *models.Link) error
	GetShortURL(ctx context.Context, link *models.Link) error
	GetStat(ctx context.Context, link *models.Link) error
}

type Repository interface {
	CreateURL(ctx context.Context, link *models.Link) error
	IncrementStat(ctx context.Context, link *models.Link) (int, error)
	FindURL(ctx context.Context, link *models.Link) error
}
