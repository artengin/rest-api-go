package postgres

import (
	"context"
	"errors"
	"github.com/artengin/rest-api-go/internal/app"
	"github.com/gocraft/dbr/v2"
	"time"
)

type PersonRepository struct {
	session *dbr.Session
}

func NewPersonRepository(session *dbr.Session) *PersonRepository {
	return &PersonRepository{session: session}
}

func (r *PersonRepository) Create(ctx context.Context, p *app.Person) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	err := r.session.
		InsertInto("person").
		Columns("email", "phone", "first_name", "last_name", "created_at", "updated_at").
		Record(p).
		Returning("id").
		LoadContext(ctx, &p.ID)
	return err
}

func (r *PersonRepository) GetAll(ctx context.Context, limit, offset int, search string) ([]*app.Person, error) {
	var persons []*app.Person
	q := r.session.Select("*").From("person")
	if search != "" {
		q = q.Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	if limit > 0 {
		q = q.Limit(uint64(limit))
	}
	if offset > 0 {
		q = q.Offset(uint64(offset))
	}
	_, err := q.LoadContext(ctx, &persons)
	return persons, err
}

func (r *PersonRepository) GetByID(ctx context.Context, id int64) (*app.Person, error) {
	var person app.Person
	err := r.session.Select("*").From("person").Where("id = ?", id).LoadOneContext(ctx, &person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *PersonRepository) Update(ctx context.Context, id int64, p *app.Person) error {
	var exists bool
	err := r.session.SelectBySql("SELECT EXISTS(SELECT 1 FROM person WHERE id = ?)", id).LoadOneContext(ctx, &exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Person not found")
	}
	p.UpdatedAt = time.Now()
	errUpdate := r.session.Update("person").
		Set("email", p.Email).
		Set("phone", p.Phone).
		Set("first_name", p.FirstName).
		Set("last_name", p.LastName).
		Set("updated_at", p.UpdatedAt).
		Where("id = ?", id).
		Returning("id").
		LoadContext(ctx, &p.ID)
	return errUpdate
}

func (r *PersonRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.session.DeleteFrom("person").Where("id = ?", id).ExecContext(ctx)
	return err
}
