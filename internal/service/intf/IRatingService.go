package intf

import (
	"app/internal/service/core/models"
	"context"

	"github.com/google/uuid"
)

type IRatingService interface {
	Create(ctx context.Context, rating *models.RatingModel) error
	GetByBookID(ctx context.Context, bookID uuid.UUID, limit, offset int) ([]*models.RatingModel, error)
	GetAvgRatingByBookID(ctx context.Context, bookID uuid.UUID) (float32, error)
}
