package intf

import (
	"app/internal/service/core/models"
	"context"

	"github.com/google/uuid"
)

type IReaderService interface {
	SignUp(ctx context.Context, reader *models.ReaderModel) error
	SignIn(ctx context.Context, phoneNumber, password string) (*models.Tokens, error)
	GetByID(ctx context.Context, ID uuid.UUID) (*models.ReaderModel, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.ReaderModel, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*models.Tokens, error)
	AddToFavorites(ctx context.Context, readerID, bookID uuid.UUID) error
}
