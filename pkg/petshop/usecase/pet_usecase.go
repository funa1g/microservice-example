package usecase

import (
	"context"

	"github.com/funa1g/microservice-example/pkg/petshop/domain"
)

type petUsecase struct {
	petRepo domain.PetRepository
	tagRepo domain.TagRepository
}

func NewPetUsecase(p domain.PetRepository, t domain.TagRepository) domain.PetUsecase {
	return &petUsecase{
		petRepo: p,
		tagRepo: t,
	}
}

func (pu *petUsecase) GetList(ctx context.Context, limit int) ([]domain.Pet, error) {
	if limit == 0 {
		limit = 20 // default limit value
	}
	return pu.petRepo.GetList(ctx, limit)
}

func (pu *petUsecase) Store(ctx context.Context, p *domain.Pet) error {
	tag, err := pu.tagRepo.GetByName(ctx, p.Tag)
	if err != nil {
		tag := &domain.Tag{
			Name: p.Name,
		}
		err = pu.tagRepo.Store(ctx, tag)
		if err != nil {
			return domain.ErrBadParamInput
		}
	}
	p.TagId = tag.ID

	return pu.petRepo.Store(ctx, p)
}

func (pu *petUsecase) GetById(ctx context.Context, id int) (domain.Pet, error) {
	return pu.petRepo.GetById(ctx, id)
}
