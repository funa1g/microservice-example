package domain

import "context"

type Pet struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Tag   string `json:"tag"`
	TagId int    `json:"tag_id"`
}

type PetUsecase interface {
	GetList(ctx context.Context, limit int) ([]Pet, error)
	Store(ctx context.Context, p *Pet) error
	GetById(ctx context.Context, id int) (Pet, error)
}

type PetRepository interface {
	GetList(ctx context.Context, limit int) ([]Pet, error)
	Store(ctx context.Context, p *Pet) error
	GetById(ctx context.Context, id int) (Pet, error)
}
