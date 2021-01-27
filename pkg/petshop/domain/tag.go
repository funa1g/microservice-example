package domain

import "context"

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}

type TagRepository interface {
	Store(ctx context.Context, t *Tag) error
	GetByName(ctx context.Context, name string) (Tag, error)
}
