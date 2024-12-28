package repository

import (
	"Marketplace-backend/internal/entity"
)

type TokenRepository interface {
	FindByToken(token string) (*entity.Token, error)
	Create(token *entity.Token) error
}
