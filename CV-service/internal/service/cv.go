package service

import (
	"CV-service/internal/core"
	"CV-service/proto"
	"context"
	"fmt"
)

type CVRepository interface {
	GetAll(ctx context.Context) ([]*core.CV, error)
}

type CVService struct {
	proto.CVServiceServer
	CVRepository CVRepository
}

func NewCVService(CVRepository CVRepository) *CVService {
	return &CVService{CVRepository: CVRepository}
}

func (s *CVService) GetAll(ctx context.Context, _ *proto.GetAllRequest) (*proto.GetAllResponse, error) {
	CVs, err := s.CVRepository.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("get all CVs in service: %w", err)
	}

	if CVs == nil {
		CVs = make([]*core.CV, 0)
	}

	CVResponses := make([]*proto.CVResponse, len(CVs))
	for i, CV := range CVs {
		CVResponses[i] = &proto.CVResponse{
			FirstName:  CV.FirstName,
			LastName:   CV.LastName,
			City:       CV.City,
			Experience: CV.Experience,
			Skills:     CV.Skills,
		}
	}

	return &proto.GetAllResponse{
		Cvs: CVResponses,
	}, nil
}
