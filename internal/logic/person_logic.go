package logic

import (
	"context"
	"github.com/artengin/rest-api-go/internal/app"
	"github.com/artengin/rest-api-go/internal/repository"
)

type PersonLogic struct {
	repo repository.PersonRepository
}

func NewPersonLogic(repo repository.PersonRepository) *PersonLogic {
	return &PersonLogic{repo: repo}
}

func (l *PersonLogic) CreatePerson(ctx context.Context, p *app.Person) error {
	return l.repo.Create(ctx, p)
}

func (l *PersonLogic) GetAllPerson(ctx context.Context, limit, offset int, search string) ([]*app.Person, error) {
	return l.repo.GetAll(ctx, limit, offset, search)
}

func (l *PersonLogic) GetPerson(ctx context.Context, id int64) (*app.Person, error) {
	return l.repo.GetByID(ctx, id)
}

func (l *PersonLogic) UpdatePerson(ctx context.Context, id int64, p *app.Person) error {
	return l.repo.Update(ctx, id, p)
}

func (l *PersonLogic) DeletePerson(ctx context.Context, id int64) error {
	return l.repo.Delete(ctx, id)
}
