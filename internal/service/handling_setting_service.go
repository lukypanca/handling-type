package service

import (
	"context"
	"tipe-handling/internal/model"
	"tipe-handling/internal/repository"
)

type HandlingSettingService struct {
	amRepo *repository.HandlingSettingRepository
}

func NewHandlingSettingService(repo *repository.HandlingSettingRepository) *HandlingSettingService {
	return &HandlingSettingService{
		amRepo: repo,
	}
}

func (s *HandlingSettingService) GetAll(ctx context.Context) ([]model.HandlingSetting, error) {
	return s.amRepo.FindAll(ctx)
}

func (s *HandlingSettingService) Create(ctx context.Context, req *model.HandlingSetting) (*model.HandlingSetting, error) {
	id, err := s.amRepo.Save(ctx, req)
	if err != nil {
		return nil, err
	}

	req.ID = id
	return req, nil
}


