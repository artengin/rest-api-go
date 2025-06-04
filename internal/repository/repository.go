package repository

import (
	"context"
	"github.com/artengin/rest-api-go/internal/app"
)

type PersonRepository interface {
	Create(ctx context.Context, p *app.Person) error
	GetAll(ctx context.Context, limit, offset int, search string) ([]*app.Person, error)
	GetByID(ctx context.Context, id int64) (*app.Person, error)
	Update(ctx context.Context, id int64, p *app.Person) error
	Delete(ctx context.Context, id int64) error
}
