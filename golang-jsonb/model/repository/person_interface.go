package repository

import (
	"context"
	"golang-jsonb/model/entity"
)

type PersonInterface interface {
	GetAllData(ctx context.Context) ([]entity.Person, error)
	GetData(ctx context.Context, name string) (entity.Person, error)
	CreateData(ctx context.Context, person entity.Person) (entity.Person, error)
	UpdateData(ctx context.Context, name string, age int) (entity.Person, error)
	DeleteData(ctx context.Context, name string) error
}
