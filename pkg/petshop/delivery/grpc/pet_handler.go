package grpc

import (
	"context"

	pb "github.com/funa1g/microservice-example/api/petshop"
	"github.com/funa1g/microservice-example/pkg/petshop/domain"
)

type PetHandler struct {
	pb.UnimplementedPetshopServer
	PUsecase domain.PetUsecase
}

func (p *PetHandler) GetPetList(ctx context.Context, req *pb.GetPetListRequest) (*pb.PetListResponse, error) {
	pets, err := p.PUsecase.GetList(ctx, int(req.GetLimit()))
	if err != nil {
		return nil, err
	}

	res := &pb.PetListResponse{}
	for _, pet := range pets {
		pr := convertPetToPetResponse(pet)
		res.Results = append(res.Results, pr)
	}
	return res, nil
}

func (p *PetHandler) PostPet(ctx context.Context, req *pb.PostPetRequest) (*pb.PetResponse, error) {
	pet := &domain.Pet{
		Name: req.GetName(),
		Tag:  req.GetTag(),
	}
	res, err := p.PUsecase.Store(ctx, pet)
	if err != nil {
		return nil, err
	}

	result := convertPetToPetResponse(res)
	return result, nil
}

func (p *PetHandler) GetPetById(ctx context.Context, req *pb.GetPetRequest) (*pb.PetResponse, error) {
	pet, err := p.PUsecase.GetById(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	result := convertPetToPetResponse(pet)

	return result, nil
}

func convertPetToPetResponse(pet domain.Pet) *pb.PetResponse {
	return &pb.PetResponse{
		Id:   int32(pet.ID),
		Name: pet.Name,
		Tag:  pet.Tag,
	}
}
