package internal

import (
	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(request CreateRequest) (CreateResponse, error) {
	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType}
	s.repository.Create(membership)
	return CreateResponse{
		ID:             membership.ID,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) GetByID(id string) (GetResponse, error) {
	membership, err := s.repository.GetById(id)
	if err != nil {
		return GetResponse{}, nil
	}
	return GetResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) deleteById(id string) error {
	err := s.repository.deleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) updateById(request UpdateRequest) (UpdateResponse, error) {
	membership, err := s.repository.UpdateById(Membership{ID: request.ID, UserName: request.UserName, MembershipType: request.MembershipType})
	if err != nil {
		return UpdateResponse{}, err
	}
	return UpdateResponse{membership.ID, membership.UserName, membership.MembershipType}, nil
}
