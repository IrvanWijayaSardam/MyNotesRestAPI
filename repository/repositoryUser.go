package repository

import (
	"context"

	"github.com/IrvanWijayaSardam/GOData/models"
)

type UserRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.User, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	Create(ctx context.Context, p *models.User) (int64, error)
	Update(ctx context.Context, p *models.User) (*models.User, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
